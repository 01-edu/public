package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"

	table "github.com/tatsushid/go-prettytable"
)

type strBoolMap map[string]bool

// Implementation of the flag.Value interface
func (a *strBoolMap) String() (res string) {
	for k, _ := range *a {
		res += k
	}
	return res
}

// Implementation of the flag.Value interface
func (a *strBoolMap) Set(str string) error {
	if *a == nil {
		*a = make(map[string]bool)
	}
	s := strings.Split(str, ",")
	for _, v := range s {
		(*a)[v] = true
	}
	return nil
}

// Flag that groups a boolean value and a regular expression
type regexpFlag struct {
	active bool
	reg    *regexp.Regexp
}

// Implementation of the flag.Value interface
func (r *regexpFlag) String() string {
	if r.reg != nil {
		return r.reg.String()
	}
	return ""
}

// Implementation of the flag.Value interface
func (r *regexpFlag) Set(s string) error {
	r.active = true
	r.reg = regexp.MustCompile(s)
	return nil
}

var (
	allowedFun = make(map[string]map[string]bool)
	allowedRep = make(map[string]int)
	// Flags
	noArrays          bool
	noSlices          bool
	noRelativeImports bool
	noTheseSlices     strBoolMap
	casting           bool
	noFor             bool
	noLit             regexpFlag
	allowBuiltin      bool
)

type illegal struct {
	T    string
	Name string
	Pos  string
}

func (i *illegal) String() string {
	return i.T + " " + i.Name + " " + i.Pos
}

func init() {
	flag.Var(&noTheseSlices, "no-these-slices", "Disallowes the slice types passed in the flag as a comma-separated list without spaces\nLike so: -no-these-slices=int,string,bool")
	flag.Var(&noLit, "no-lit",
		`The use of basic literals (strings or characters) matching the pattern -no-lit="{PATTERN}"
passed to the program would not be allowed`,
	)
	flag.BoolVar(&noRelativeImports, "no-relative-imports", false, `Disallowes the use of relative imports`)
	flag.BoolVar(&noFor, "no-for", false, `The "for" instruction is not allowed`)
	flag.BoolVar(&casting, "cast", false, "Allowes casting")
	flag.BoolVar(&noArrays, "no-array", false, "Deprecated: use -no-slices")
	flag.BoolVar(&noSlices, "no-slices", false, "Disallowes all slice types")
	flag.BoolVar(&allowBuiltin, "allow-builtin", false, "Allowes all builtin functions and casting")
	sort.Sort(sort.StringSlice(os.Args[1:]))
}

func main() {
	flag.Parse()
	filename := goFile(flag.Args())

	if _, err := os.Stat(filename); err != nil {
		fmt.Printf("\t%s\n", err)
		os.Exit(1)
	}

	err := parseArgs(flag.Args())

	if err != nil {
		fmt.Printf("\t%s\n", err)
		os.Exit(1)
	}

	load := make(loadedSource)

	currentPath := filepath.Dir(filename)

	err = loadProgram(currentPath, load)

	if err != nil {
		fmt.Printf("\t%s\n", err)
		os.Exit(1)
	}

	info := analyzeProgram(filename, currentPath, load)

	if info.illegals != nil {
		fmt.Println("Cheating:")
		printIllegals(info.illegals)
		os.Exit(1)
	}
}

func goFile(args []string) string {
	for _, v := range args {
		if strings.HasSuffix(v, ".go") {
			return v
		}
	}
	return ""
}

// Returns the smallest block containing the position pos. It can
// return nil if `pos` is not inside any ast.BlockStmt
func smallestBlock(pos token.Pos, blocks []*ast.BlockStmt) (minBlock *ast.BlockStmt) {
	var minSize token.Pos
	for _, v := range blocks {
		if pos > v.Pos() && pos < v.End() {
			size := v.End() - v.Pos()
			if minBlock == nil || size < minSize {
				minBlock = v
				minSize = size
			}
		}
	}
	return minBlock
}

// Used to mark an ast.Object as a function parameter
type data struct {
	parameter bool
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

// Returns true if `block` is contained inside another ast.BlockStmt
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
	scopes := make(map[*ast.BlockStmt]*ast.Scope)
	if l.blocks == nil {
		return nil
	}
	for _, b := range l.blocks {
		if !isContained(b, l.blocks) {
			scopes[b] = ast.NewScope(pkgScope)
		}
	}
	for _, b := range l.blocks {
		if scopes[b] == nil {
			createChildScope(b, l, scopes)
		}
	}
	return scopes
}

type blockVisitor struct {
	funct []*function
	// All functions defined in the scope in any
	// way: as a funcDecl, GenDecl or AssigmentStmt
	oneBlock bool
	// Indicates if the visitor already encounter a
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
		def := extractFunction(t)
		if def == nil || def.obj == nil {
			return b
		}
		b.funct = append(b.funct, def)
		return nil
	default:
		return b
	}
}

type loadedSource map[string]*loadVisitor

// Returns information about the function defined in the block node
func functionsInfo(block ast.Node) []*function {
	b := &blockVisitor{}
	ast.Walk(b, block)
	return b.funct
}

func (l *loadVisitor) set() {
	l.functions = make(map[string]ast.Node)
	l.absImports = make(map[string]*element)
	l.relImports = make(map[string]*element)
	l.objFunc = make(map[*ast.Object]ast.Node)
	l.fset = token.NewFileSet()
	l.scopes = make(map[*ast.BlockStmt]*ast.Scope)

}

func loadProgram(path string, load loadedSource) error {
	l := &loadVisitor{}
	l.set()

	pkgs, err := parser.ParseDir(l.fset, path, nil, parser.AllErrors)

	if err != nil {
		return err
	}

	if len(pkgs) > 1 {
		packages := []string{}
		for pkgName := range pkgs {
			packages = append(packages, pkgName)
		}
		return fmt.Errorf("There should be only one package in this directory: Found packages: %v", packages)
	}

	for _, pkg := range pkgs {
		ast.Walk(l, pkg)
		l.pkgScope = ast.NewScope(nil)
		functions := functionsInfo(pkg)
		for _, f := range functions {
			l.pkgScope.Insert(f.obj)
		}
		l.scopes = createScopes(l, l.pkgScope)
		fillScope(functions, l.pkgScope, l.scopes)
		for block, scope := range l.scopes {
			functions := functionsInfo(block)
			fillScope(functions, scope, l.scopes)
		}
		load[path] = l
		l.files = pkg.Files
	}

	for _, relativePath := range l.relImports {
		if load[relativePath.name] == nil {
			newPath := filepath.Clean(path + "/" + relativePath.name)
			err = loadProgram(newPath, load)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func smallestScopeContaining(pos token.Pos, path string, load loadedSource) *ast.Scope {
	pack := load[path]
	sm := smallestBlock(pos, pack.blocks)
	if sm == nil {
		return pack.pkgScope
	}

	return pack.scopes[sm]
}

func lookupDefinitionObj(el *element, path string, load loadedSource) *ast.Object {
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
	fset           *token.FileSet
	uses           []*element
	selections     map[string][]*element
	arrays         []*occurrence
	lits           []*occurrence
	fors           []*occurrence
	callRepetition map[string]int
	oneTime        bool
}

func (v *visitor) getPos(n ast.Node) string {
	return v.fset.Position(n.Pos()).String()
}

func (v *visitor) Visit(n ast.Node) ast.Visitor {
	switch t := n.(type) {
	case *ast.FuncDecl, *ast.GenDecl, *ast.AssignStmt:
		//Avoids analyzing a nested declarations
		//Since this is handle by the functions `isAllowed`
		fdef := extractFunction(t)
		if fdef == nil || fdef.obj == nil {
			return v
		}
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
			v.uses = append(v.uses, &element{
				name: fun.Name,
				pos:  fun.Pos(),
			})
			v.callRepetition[fun.Name]++
		}

	case *ast.SelectorExpr:
		if x, ok := t.X.(*ast.Ident); ok {
			v.selections[x.Name] = append(v.selections[x.Name], &element{
				name: t.Sel.Name,
				pos:  n.Pos(),
			})
			v.callRepetition[x.Name+"."+t.Sel.Name]++
		}
	}
	return v
}

func (v *visitor) set(fset *token.FileSet) {
	v.selections = make(map[string][]*element)
	v.callRepetition = make(map[string]int)
	v.fset = fset
}

func (info *info) add(v *visitor) {
	info.fors = append(info.fors, v.fors...)
	info.lits = append(info.lits, v.lits...)
	info.arrays = append(info.arrays, v.arrays...)
	for name, v := range v.callRepetition {
		info.callRepetition[name] += v
	}
}

// Returns the info structure with all the ocurrences of the element
// of the analised in the project
// TODO: Refactor so this function has only one responsibility
func isAllowed(function *element, path string, load loadedSource, walked map[ast.Node]bool, info *info) bool {
	functionObj := lookupDefinitionObj(function, path, load)
	definedLocally := functionObj != nil
	explicitlyAllowed := allowedFun["builtin"]["*"] || allowedFun["builtin"][function.name]

	isFunctionParameter := func(function *ast.Object) bool {
		arg, ok := function.Data.(data)
		return ok && arg.parameter
	}

	DoesntCallMoreFunctions := func(functionDefinition ast.Node, v *visitor) bool {
		if !walked[functionDefinition] {
			ast.Walk(v, functionDefinition)
			info.add(v)
			walked[functionDefinition] = true
		}

		return v.uses == nil && v.selections == nil
	}

	appendIllegalCall := func(function *element) {
		info.illegals = append(info.illegals, &illegal{
			T:    "illegal-call",
			Name: function.name,
			Pos:  load[path].fset.Position(function.pos).String(),
		})

	}
	if !definedLocally && !explicitlyAllowed {
		appendIllegalCall(function)
		return false
	}

	functionDefinition := load[path].objFunc[functionObj]
	v := &visitor{}
	v.set(load[path].fset)

	if explicitlyAllowed || isFunctionParameter(functionObj) ||
		DoesntCallMoreFunctions(functionDefinition, v) {
		return true
	}

	allowed := true
	for _, functionCall := range v.uses {
		if !isAllowed(functionCall, path, load, walked, info) {
			appendIllegalCall(functionCall)
			allowed = false
		}
	}

	for pck, funcNames := range v.selections {
		pathToFunction := func() string { return load[path].relImports[pck].name }
		isRelativeImport := load[path].relImports[pck] != nil
		for _, fun := range funcNames {
			appendIllegalAccess := func() {
				info.illegals = append(info.illegals, &illegal{
					T:    "illegal-access",
					Name: pck + "." + fun.name,
					Pos:  load[path].fset.Position(fun.pos).String(),
				})
				allowed = false
			}

			absoluteImport := load[path].absImports[pck]
			importExplicitlyAllowed := absoluteImport == nil ||
				allowedFun[absoluteImport.name][fun.name] ||
				allowedFun[absoluteImport.name]["*"]

			if !isRelativeImport && !importExplicitlyAllowed {
				appendIllegalAccess()
			} else if isRelativeImport &&
				!isAllowed(newElement(fun.name), filepath.Clean(path+"/"+pathToFunction()), load, walked, info) {
				appendIllegalAccess()
			}
		}
	}

	if !allowed {
		info.illegals = append(info.illegals, &illegal{
			T:    "illegal-definition",
			Name: functionObj.Name,
			Pos:  load[path].fset.Position(functionDefinition.Pos()).String(),
		})
	}
	return allowed
}

func removeRepetitions(slc []*illegal) (result []*illegal) {
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
	arrays         []*occurrence
	lits           []*occurrence
	fors           []*occurrence
	callRepetition map[string]int
	illegals       []*illegal // functions, selections that are not allowed
}

func newElement(name string) *element {
	return &element{
		name: name,
		pos:  token.Pos(0),
	}

}

func analyzeProgram(filename, path string, load loadedSource) *info {
	fset := load[path].fset
	file := load[path].files[filename]
	functions := functionsInfo(file)

	info := &info{
		callRepetition: make(map[string]int),
	}

	info.illegals = append(info.illegals, analyzeImports(file, fset, noRelativeImports)...)

	walked := make(map[ast.Node]bool)

	for _, fun := range functions {
		function := newElement(fun.obj.Name)
		isAllowed(function, path, load, walked, info)
	}

	info.illegals = append(info.illegals, analyzeLoops(info.fors, noFor)...)
	info.illegals = append(info.illegals, analyzeArrayTypes(info.arrays, noArrays || noSlices, noTheseSlices)...)
	info.illegals = append(info.illegals, analyzeLits(info.lits, noLit)...)
	info.illegals = append(info.illegals, analyzeRepetition(info.callRepetition, allowedRep)...)
	info.illegals = removeRepetitions(info.illegals)
	return info
}

func parseArgs(toAllow []string) error {
	allowBuiltins()
	allowCasting()
	for _, v := range toAllow {
		err := allowFunction(v)
		if err != nil {
			return err
		}
	}
	return nil
}

func allowFunction(functionPath string) error {
	functionName := functionName(functionPath)
	packageName := packageName(functionPath)
	// for github.com/01-edu/z01 shortName = z01
	packageShortName := filepath.Base(packageName)
	restrictsRepetitions := strings.ContainsRune(functionPath, '#')
	if restrictsRepetitions {
		allowedReps, err := repetitionsAllowed(functionPath)
		if err != nil {
			return err
		}
		allowedRep[packageShortName+"."+functionName] = allowedReps
	}
	if allowedFun[packageName] == nil {
		allowedFun[packageName] = make(map[string]bool)
	}
	allowedFun[packageName][functionName] = true
	return nil
}

func functionName(functionPath string) string {
	segmentedPath := strings.Split(functionPath, ".")
	return strings.Split(segmentedPath[len(segmentedPath)-1], "#")[0]
}

func packageName(functionPath string) string {
	segmentedPath := strings.Split(functionPath, ".")
	hasNoPackage := len(segmentedPath) < 2
	if hasNoPackage {
		return "builtin"
	}
	return strings.Join(segmentedPath[:len(segmentedPath)-1], ".")
}

// Assumes that `functionPath` contains `#`
func repetitionsAllowed(functionPath string) (int, error) {
	segmentedPath := strings.Split(functionPath, "#")
	repetitions := segmentedPath[len(segmentedPath)-1]
	allowedReps, err := strconv.Atoi(repetitions)
	if err != nil {
		return allowedReps, fmt.Errorf("After the '#' there should be an integer" +
			" representing the maximum number of allowed occurrences")
	}
	return allowedReps, nil
}

func allowBuiltins() {
	if allowedFun["builtin"] == nil {
		allowedFun["builtin"] = make(map[string]bool)
	}
	if allowBuiltin {
		allowedFun["builtin"]["*"] = true
	}
}

func allowCasting() {
	if allowedFun["builtin"] == nil {
		allowedFun["builtin"] = make(map[string]bool)
	}

	predeclaredTypes := []string{"bool", "byte", "complex64", "complex128",
		"error", "float32", "float64", "int", "int8",
		"int16", "int32", "int64", "rune", "string",
		"uint", "uint8", "uint16", "uint32", "uint64",
		"uintptr",
	}

	if casting {
		for _, v := range predeclaredTypes {
			allowedFun["builtin"][v] = true
		}
	}
}

func printIllegals(illegals []*illegal) {
	tbl, err := table.NewTable([]table.Column{
		{Header: "\tTYPE:"},
		{Header: "NAME:", MinWidth: 7},
		{Header: "LOCATION:"},
	}...)
	if err != nil {
		panic(err)
	}
	tbl.Separator = "\t"
	for _, v := range illegals {
		tbl.AddRow("\t"+v.T, v.Name, v.Pos)
	}
	tbl.Print()
}

func analyzeRepetition(callRepetition map[string]int, allowRep map[string]int) (illegals []*illegal) {
	for name, rep := range allowedRep {
		if callRepetition[name] > rep {
			diff := callRepetition[name] - rep
			illegals = append(illegals, &illegal{
				T:    "illegal-amount",
				Name: name + " exeding max repetitions by " + strconv.Itoa(diff),
				Pos:  "all the project",
			})
		}
	}
	return illegals
}

func analyzeLits(litOccu []*occurrence, noLit regexpFlag) (illegals []*illegal) {
	if noLit.active {
		for _, v := range litOccu {
			if noLit.reg.Match([]byte(v.name)) {
				illegals = append(illegals, &illegal{
					T:    "illegal-lit",
					Name: v.name,
					Pos:  v.pos,
				})
			}
		}
	}
	return illegals
}

func analyzeArrayTypes(arrays []*occurrence, noArrays bool, noTheseSlices map[string]bool) (illegals []*illegal) {
	for _, v := range arrays {
		if noArrays || noTheseSlices[v.name] {
			illegals = append(illegals, &illegal{
				T:    "illegal-slice",
				Name: v.name,
				Pos:  v.pos,
			})
		}
	}
	return illegals
}

func analyzeLoops(fors []*occurrence, noFor bool) (illegals []*illegal) {
	if noFor {
		for _, v := range fors {
			illegals = append(illegals, &illegal{
				T:    "illegal-loop",
				Name: v.name,
				Pos:  v.pos,
			})
		}
	}
	return illegals

}

type importVisitor struct {
	imports map[string]*element
}

func (i *importVisitor) Visit(n ast.Node) ast.Visitor {
	if imp, ok := n.(*ast.ImportSpec); ok {
		path, _ := strconv.Unquote(imp.Path.Value)
		var name string
		if imp.Name != nil {
			name = imp.Name.Name
		} else {
			name = filepath.Base(path)
		}
		el := &element{
			name: path,
			pos:  n.Pos(),
		}
		i.imports[name] = el
	}
	return i
}

func analyzeImports(file ast.Node, fset *token.FileSet, noRelImp bool) (illegals []*illegal) {
	i := &importVisitor{
		imports: make(map[string]*element),
	}
	ast.Walk(i, file)
	for _, path := range i.imports {
		isRelativeImport := isRelativeImport(path.name)
		if (noRelativeImports && isRelativeImport) || (allowedFun[path.name] == nil && !isRelativeImport) {
			illegals = append(illegals, &illegal{
				T:    "illegal-import",
				Name: path.name,
				Pos:  fset.Position(path.pos).String(),
			})
		}
	}
	return illegals
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
	scopes     map[*ast.BlockStmt]*ast.Scope
	// nil after the visit
	// used to keep the result of the createScope function
	pkgScope *ast.Scope
	files    map[string]*ast.File
}

func (l *loadVisitor) String() (res string) {
	res = "files"
	for f, _ := range l.files {
		res += f + ","
	}
	return res
}

// Returns all the parameter of a function that identify a function
func functionsInTheParameters(params *ast.FieldList) []string {
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

type function struct {
	obj    *ast.Object // the ast.Object that represents a function
	params []string
	// the name of the parameter that represent
	// functions
	body *ast.BlockStmt
}

// Returns information about a node representing a function declaration
func extractFunction(n ast.Node) *function {
	function := &function{}
	switch t := n.(type) {
	case *ast.FuncDecl:
		function.obj = t.Name.Obj
		function.params = functionsInTheParameters(t.Type.Params)
		function.body = t.Body
		return function
	case *ast.GenDecl:
		for _, v := range t.Specs {
			if val, ok := v.(*ast.ValueSpec); ok {
				for i, value := range val.Values {
					if funcLit, ok := value.(*ast.FuncLit); ok {
						function.obj = val.Names[i].Obj
						function.params = functionsInTheParameters(funcLit.Type.Params)
						function.body = funcLit.Body
					}
				}
			}
		}
		return function
	case *ast.AssignStmt:
		for i, right := range t.Rhs {
			if funcLit, ok := right.(*ast.FuncLit); ok {
				if ident, ok := t.Lhs[i].(*ast.Ident); ok {
					function.obj = ident.Obj
					function.params = functionsInTheParameters(funcLit.Type.Params)
				}
			}
			return function
		}
	default:
		return function
	}
	return function
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
		} else {
			l.absImports[name] = el
		}
	case *ast.FuncDecl, *ast.GenDecl, *ast.AssignStmt:
		fdef := extractFunction(t)
		if fdef == nil || fdef.obj == nil {
			return l
		}
		l.objFunc[fdef.obj] = n
	case *ast.BlockStmt:
		l.blocks = append(l.blocks, t)
	}
	return l
}

// Returns true if the string matches the format of a relative import
func isRelativeImport(s string) bool {
	return strings.HasPrefix(s, ".")
}
