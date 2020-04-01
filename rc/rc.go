package main

import (
	"flag"
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"

type strBoolMap map[string]bool

func (a *strBoolMap) String() string {
	var res string
	for k, _ := range *a {
		res += k
	}
	return res
}

func (a *strBoolMap) Set(str string) error {
	if *a == nil {
		*a = make(map[string]bool)
	}
	s := strings.Split(str, " ")
	for _, v := range s {
		(*a)[v] = true
	}
	return nil
}

type arrFlag struct {
	active  bool
	content []string
}

func (a *arrFlag) String() string {
	return strings.Join(a.content, " ")
}

func (a *arrFlag) Set(s string) error {
	a.active = true
	a.content = strings.Split(s, ",")
	return nil
}

// flag that groups a boolean value and a regular expression
type regexpFlag struct {
	active bool
	reg    *regexp.Regexp
}

func (r *regexpFlag) String() string {
	if r.reg != nil {
		return r.reg.String()
	}
	return ""
}

func (r *regexpFlag) Set(s string) error {
	re := regexp.MustCompile(s)
	r.active = true
	r.reg = re
	return nil
}

var (
	allowedFun = make(map[string]map[string]bool)
	allowedRep = make(map[string]int)
	// Flags
	noArrays          bool
	noRelativeImports bool
	noTheseArrays     strBoolMap
	casting           bool
	noFor             bool
	noLit             regexpFlag
	allowBuiltin      bool
)

//pkgFunc for all the functions of a given package
type pkgFunc struct {
	functions []string
	path      string
}

type funcImp struct {
	pkg, fun string
	pos      token.Pos
}

// All visitors
type callVisitor struct {
	Calls []string
	Fset  *token.FileSet
}

// Get the position of the node in the file
type locate interface {
	getPos(ast.Node) string
}

type illegal struct {
	T    string
	Name string
	Pos  string
}

func (i *illegal) String() string {
	return i.T + " " + i.Name + " " + i.Pos
}

// Returns the smallest block containing the position pos. It can
// return nil if `pos` is not inside any ast.BlockStmt
func smallestBlock(pos token.Pos, blocks []*ast.BlockStmt) *ast.BlockStmt {
	var minBlk *ast.BlockStmt
	var minSize token.Pos
	for _, v := range blocks {
		if pos > v.Pos() && pos < v.End() {
			size := v.End() - v.Pos()
			if minBlk == nil || size < minSize {
				minBlk = v
				minSize = size
			}
		}
	}
	return minBlk
}

type data struct {
	argument bool
}

func fillScope(funcDefs []*fDefInfo, scope *ast.Scope, scopes map[*ast.BlockStmt]*ast.Scope) {
	for _, fun := range funcDefs {
		scope.Insert(fun.obj)
		for _, name := range fun.paramsFunc {
			obj := ast.NewObj(ast.Fun, name)

			data := data{
				argument: true,
			}
			obj.Data = data
			scopes[fun.body].Insert(obj)
		}
	}
}

// Create the scopes for a BlockStmt contained inside another BlockStmt
func createChildScope(block *ast.BlockStmt, l *loadVisitor, scopes map[*ast.BlockStmt]*ast.Scope) {
	blocks := l.blocks
	// The smalles block containing the beggining of the block
	parentBlock := smallestBlock(block.Pos(), blocks)
	if scopes[parentBlock] == nil {
		createChildScope(parentBlock, l, scopes)
	}
	scopes[block] = ast.NewScope(scopes[parentBlock])
}

// Returns true `block` is contained inside another block
func isContained(block *ast.BlockStmt, blocks []*ast.BlockStmt) bool {
	for _, v := range blocks {
		if block == v {
			continue
		}
		if block.Pos() > v.Pos() && block.End() < v.End() {
			return true
		}
	}
	return false
}

// Creates all the scopes in the package
func createScopes(l *loadVisitor, pkgScope *ast.Scope) map[*ast.BlockStmt]*ast.Scope {
	blocks := l.blocks
	scopes := make(map[*ast.BlockStmt]*ast.Scope)
	if blocks == nil {
		return nil
	}
	for _, b := range blocks {
		if !isContained(b, blocks) {
			scopes[b] = ast.NewScope(pkgScope)
			continue
		}
	}
	for _, b := range blocks {
		if scopes[b] != nil {
			continue
		}
		createChildScope(b, l, scopes)
	}
	return scopes
}

type blockVisitor struct {
	fdef []*fDefInfo // All functions defined in the scope in any
	// way: as a funcDecl, GenDecl or AssigmentStmt
	oneBlock bool // Indicates if the visitor already encounter a
	// blockStmt
}

func (b *blockVisitor) Visit(n ast.Node) ast.Visitor {
	switch t := n.(type) {
	case *ast.BlockStmt:
		if b.oneBlock {
			return nil
		}
		return b
	case *ast.FuncDecl, *ast.GenDecl, *ast.AssignStmt:
		def := funcInfo(t)
		if def == nil || def.obj == nil {
			return b
		}
		b.fdef = append(b.fdef, def)
		return nil
	default:
		return b
	}
}

type loadedSource map[string]*loadVisitor

// Returns information about the function defined in the block node
func defs(block ast.Node) []*fDefInfo {
	b := &blockVisitor{}
	ast.Walk(b, block)
	return b.fdef
}

func loadProgram(path string, load loadedSource) error {
	l := &loadVisitor{
		functions:  make(map[string]ast.Node),
		absImports: make(map[string]*element),
		relImports: make(map[string]*element),
		objFunc:    make(map[*ast.Object]ast.Node),
		fset:       token.NewFileSet(),
		scopes:     make(map[*ast.BlockStmt]*ast.Scope),
	}

	pkgs, err := parser.ParseDir(l.fset, path, nil, parser.AllErrors)

	if err != nil {
		return err
	}

	for _, pkg := range pkgs {
		ast.Walk(l, pkg)
		l.pkgScope = ast.NewScope(nil)
		def := defs(pkg)
		for _, v := range def {
			l.pkgScope.Insert(v.obj)
		}
		l.scopes = createScopes(l, l.pkgScope)
		fillScope(def, l.pkgScope, l.scopes)
		for block, scope := range l.scopes {
			defs := defs(block)
			fillScope(defs, scope, l.scopes)
		}
		load[path] = l
	}

	for _, v := range l.relImports {
		if load[v.name] == nil {
			newPath, _ := filepath.Abs(path + "/" + v.name)
			err = loadProgram(newPath, load)
			if err != nil {
				return err
			}
		}
	}
	return err
}

func smallestScopeContaining(pos token.Pos, path string, load loadedSource) *ast.Scope {
	pack := load[path]
	sm := smallestBlock(pos, pack.blocks)
	if sm == nil {
		return pack.pkgScope
	}

	return pack.scopes[sm]
}

func lookupDefinitionObj(el element, path string, load loadedSource) *ast.Object {
	scope := smallestScopeContaining(el.pos, path, load)
	for scope != nil {
		obj := scope.Lookup(el.name)
		if obj != nil {
			return obj
		}
		scope = scope.Outer
	}
	return nil
}

type visitor struct {
	fset       *token.FileSet
	uses       []element
	selections map[string][]*element
	arrays     []*occurrence
	lits       []*occurrence
	fors       []*occurrence
	callRep    map[string]int
	oneTime    bool
}

func (v *visitor) getPos(n ast.Node) string {
	return v.fset.Position(n.Pos()).String()
}

func (v *visitor) Visit(n ast.Node) ast.Visitor {
	switch t := n.(type) {
	case *ast.FuncDecl, *ast.GenDecl, *ast.AssignStmt:
		//Avoids analysing a declaration inside a declaration
		//Since this is handle by the functions `isAllowed`
		if v.oneTime {
			return nil
		}
		v.oneTime = true
		return v
	case *ast.BasicLit:
		if t.Kind != token.CHAR && t.Kind != token.STRING {
			return nil
		}
		v.lits = append(v.lits, &occurrence{pos: v.getPos(n), name: t.Value})

	case *ast.ArrayType:
		if op, ok := t.Elt.(*ast.Ident); ok {
			v.arrays = append(v.arrays, &occurrence{
				name: op.Name,
				pos:  v.getPos(n),
			})
		}
	case *ast.ForStmt:
		v.fors = append(v.fors, &occurrence{
			name: "for",
			pos:  v.getPos(n),
		})
	case *ast.CallExpr:
		if fun, ok := t.Fun.(*ast.Ident); ok {
			v.uses = append(v.uses, element{
				name: fun.Name,
				pos:  fun.Pos(),
			})
			v.callRep[fun.Name]++
		}

	case *ast.SelectorExpr:
		if x, ok := t.X.(*ast.Ident); ok {
			v.selections[x.Name] = append(v.selections[x.Name], &element{
				name: t.Sel.Name,
				pos:  n.Pos(),
			})
			v.callRep[x.Name+"."+t.Sel.Name]++
		}
	}
	return v
}

// Returns the info structure with all the ocurrences of the element
// of the analised in the project
func isAllowed(function element, path string, load loadedSource, walked map[ast.Node]bool, info *info) bool {
	if walked == nil {
		walked = make(map[ast.Node]bool)
	}
	fdef := lookupDefinitionObj(function, path, load)
	if fdef == nil && !allowedFun["builtin"]["*"] && !allowedFun["builtin"][function.name] {
		info.illegals = append(info.illegals, &illegal{
			T:    "illegal-call",
			Name: function.name,
			Pos:  load[path].fset.Position(function.pos).String(),
		})
		return false
	}
	if fdef == nil {
		return true
	}
	if arg, ok := fdef.Data.(data); ok && arg.argument {
		return true
	}
	funcNode := load[path].objFunc[fdef]
	v := &visitor{
		selections: make(map[string][]*element),
		callRep:    make(map[string]int),
		fset:       load[path].fset,
	}
	if !walked[funcNode] {
		ast.Walk(v, funcNode)
		info.fors = append(info.fors, v.fors...)
		info.lits = append(info.lits, v.lits...)
		info.arrays = append(info.arrays, v.arrays...)
		for name, v := range v.callRep {
			info.callRep[name] += v
		}
		walked[funcNode] = true
	}

	if v.uses == nil && v.selections == nil {
		return true
	}

	allowed := true
	for _, use := range v.uses {
		allowedUse := isAllowed(use, path, load, walked, info)
		if !allowedUse {
			info.illegals = append(info.illegals, &illegal{
				T:    "illegal-call",
				Name: use.name,
				Pos:  load[path].fset.Position(use.pos).String(),
			})
		}
		allowed = allowedUse && allowed
	}

	for pck, funcNames := range v.selections {
		importRelPath := load[path].relImports[pck]
		for _, fun := range funcNames {
			if importRelPath == nil {
				absImp := load[path].absImports[pck]
				if absImp != nil && !allowedFun[absImp.name][fun.name] && !allowedFun[absImp.name]["*"] {
					// Add to the illegals array the import and selection
					info.illegals = append(info.illegals, &illegal{
						T:    "illegal-access",
						Name: pck + "." + fun.name,
						Pos:  load[path].fset.Position(fun.pos).String(),
					})
					allowed = false
				}
				continue
			}

			newPath, err := filepath.Abs(path + "/" + importRelPath.name)

			if err != nil {
				panic(err)
			}
			newEl := element{
				name: fun.name,
				pos:  token.Pos(0),
			}
			allowedSel := isAllowed(newEl, newPath, load, walked, info)
			if !allowedSel {
				info.illegals = append(info.illegals, &illegal{
					T:    "illegal-access",
					Name: pck + "." + fun.name,
					Pos:  load[path].fset.Position(fun.pos).String(),
				})
			}
			allowed = allowedSel && allowed
		}
	}
	if !allowed {
		info.illegals = append(info.illegals, &illegal{
			T:    "illegal-definition",
			Name: fdef.Name,
			Pos:  load[path].fset.Position(funcNode.Pos()).String(),
		})
	}
	return allowed
}

func removeRepetitions(slc []*illegal) []*illegal {
	var result []*illegal
	in := make(map[string]bool)
	for _, v := range slc {
		if in[v.Pos] {
			continue
		}
		result = append(result, v)
		in[v.Pos] = true
	}
	return result
}

type occurrence struct {
	name string
	pos  string
}

type info struct {
	arrays   []*occurrence
	lits     []*occurrence
	fors     []*occurrence
	callRep  map[string]int
	illegals []*illegal // functions, selections that are not allowed
}

func analyseProgram(functions []*fDefInfo, path string, load loadedSource) *info {
	info := &info{
		callRep: make(map[string]int),
	}

	walked := make(map[ast.Node]bool)

	for _, v := range functions {
		f := element{
			name: v.obj.Name,
			pos:  token.Pos(0),
		}
		isAllowed(f, path, load, walked, info)
	}

	info.illegals = removeRepetitions(info.illegals)
	return info
}

//reformat from the data base
func splitArgs(args string) []string {
	result := strings.Split(args, " ")
	return result
}

func goFile(args []string) string {
	for _, v := range args {
		if strings.HasSuffix(v, ".go") {
			return v
		}
	}
	return ""
}

type flags struct {
	l struct { // flag for char or string literal
		noLit   bool   // true -> unallows
		pattern string // this pattern
	}
}

func fillScope(funcDefs []*function, scope *ast.Scope, scopes map[*ast.BlockStmt]*ast.Scope) {
	for _, fun := range funcDefs {
		scope.Insert(fun.obj)
		for _, name := range fun.params {
			obj := ast.NewObj(ast.Fun, name)
			obj.Data = data{
				parameter: true,
			}
			scopes[fun.body].Insert(obj)
		}
	}
}

// Create the scopes for a BlockStmt contained inside another BlockStmt
func createChildScope(
	block *ast.BlockStmt,
	l *loadVisitor, scopes map[*ast.BlockStmt]*ast.Scope) {
	blocks := l.blocks
	// The smalles block containing the beggining of the block
	parentBlock := smallestBlock(block.Pos(), blocks)
	if scopes[parentBlock] == nil {
		createChildScope(parentBlock, l, scopes)
	}
	scopes[block] = ast.NewScope(scopes[parentBlock])
}

func init() {
	flag.Var(&noTheseArrays, "no-these-arrays", "unallowes the array types passed in the flag")
	flag.Var(&noLit, "no-lit",
		`The use of string literals matching the pattern --no-lit="{PATTERN}"`+
			`passed to the program would not be allowed`,
	)
	flag.BoolVar(&noRelativeImports, "no-relative-imports", false, `No disallowes the use of relative imports`)
	flag.BoolVar(&noFor, "no-for", false, `The "for" instruction is not allowed`)
	flag.BoolVar(&casting, "cast", false, "allowes casting")
	flag.BoolVar(&noArrays, "no-arrays", false, "unallowes the array types passed in the flag")
	flag.BoolVar(&allowBuiltin, "allow-builtin", false, "Allowes all builtin functions and casting")
}

func main() {
}

type element struct {
	name string
	pos  token.Pos
}

type loadVisitor struct {
	relImports map[string]*element
	absImports map[string]*element
	functions  map[string]ast.Node
	fset       *token.FileSet
	objFunc    map[*ast.Object]ast.Node
	blocks     []*ast.BlockStmt
	scopes     map[*ast.BlockStmt]*ast.Scope // nil after the visit
	// used to keep the result of the createScope function
	pkgScope *ast.Scope
}

// Returns all the parameter of a function that identify a function
func listParamFunc(params *ast.FieldList) []string {
	var funcs []string
	for _, param := range params.List {
		if _, ok := param.Type.(*ast.FuncType); ok {
			for _, name := range param.Names {
				funcs = append(funcs, name.Name)
			}
		}
	}
	return funcs
}

type fDefInfo struct {
	obj        *ast.Object // the object that represents a function
	paramsFunc []string    // the name of the parameter that represent
	// functions
	body *ast.BlockStmt
}

// Returns information about a node representing a function declaration
func funcInfo(n ast.Node) *fDefInfo {
	fdef := &fDefInfo{}
	switch t := n.(type) {
	case *ast.FuncDecl:
		fdef.obj = t.Name.Obj
		fdef.paramsFunc = listParamFunc(t.Type.Params)
		fdef.body = t.Body
		return fdef
	case *ast.GenDecl:
		for _, v := range t.Specs {
			if val, ok := v.(*ast.ValueSpec); ok {
				for i, value := range val.Values {
					if funcLit, ok := value.(*ast.FuncLit); ok {
						fdef.obj = val.Names[i].Obj
						fdef.paramsFunc = listParamFunc(funcLit.Type.Params)
						fdef.body = funcLit.Body
					}
				}
			}
		}
		return fdef
	case *ast.AssignStmt:
		for i, right := range t.Rhs {
			if funcLit, ok := right.(*ast.FuncLit); ok {
				if ident, ok := t.Lhs[i].(*ast.Ident); ok {
					fdef.obj = ident.Obj
					fdef.paramsFunc = listParamFunc(funcLit.Type.Params)
				}
			}
			return fdef
		}
	default:
		return fdef
	}
	return fdef
}
func (l *loadVisitor) Visit(n ast.Node) ast.Visitor {
	switch t := n.(type) {
	case *ast.ImportSpec:
		path, _ := strconv.Unquote(t.Path.Value)
		var name string
		if t.Name != nil {
			name = t.Name.Name
		} else {
			name = filepath.Base(path)
		}
		el := &element{
			name: path,
			pos:  n.Pos(),
		}

		if isRelativeImport(path) {
			l.relImports[name] = el
			break
		}
		l.absImports[name] = el
	case *ast.FuncDecl, *ast.GenDecl, *ast.AssignStmt:
		fdef := funcInfo(t)
		if fdef == nil || fdef.obj == nil {
			return l
		}
		l.objFunc[fdef.obj] = n
	case *ast.BlockStmt:
		l.blocks = append(l.blocks, t)
	}
	return l
}

func (f flags) isLitAllowed(s string) bool {
	matched, err := regexp.Match(f.l.pattern, []byte(s))

	if err != nil {
		return true
	}

	return pack.scopes[sm]
}

// Returns true if the string matches the format of a relative import
func isRelativeImport(s string) bool {
	reg := regexp.MustCompile(`^\.`)
	return reg.Match([]byte(s))
}
