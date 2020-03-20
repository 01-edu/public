package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	identation = "  "
)

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
	a.content = strings.Split(s, " ")
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
	allowedImp map[string]map[string]bool // Map of the allowed imports
	allowedFun map[string]bool            // Map of the allowed built-in functions
	// Is necessary an array to keep all the call instances.
	callX []nodePos // Keeps the name of the called functions and the position in the file.
	// A map is enough for function declarations because they are unique.
	funcDeclPkg      map[string]*funcBody // Keeps the name of the function associated to its body and its position in the file.
	allArrayTypes    = true
	arraysInstances  []nodePos
	forStmts         []nodePos
	basicLits        []nodePos
	illegals         []illegal
	notAllowedArrayT []string
	predeclaredTypes = []string{"bool", "byte", "complex64", "complex128",
		"error", "float32", "float64", "int", "int8",
		"int16", "int32", "int64", "rune", "string",
		"uint", "uint8", "uint16", "uint32", "uint64",
		"uintptr",
	}
	relativeImports []string
	importPkg       map[string]*pkgFunc
	pkgName         []string
	allImports      map[string]bool
	openImports     []string
	funcOccurrences map[string]int
	// Flags
	noArrays          bool
	noRelativeImports bool
	noTheseArrays     strBoolMap
	casting           bool
	noFor             bool
	noLit             regexpFlag
)

// pkgFunc for all the functions of a given package
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

type fileVisitor struct {
	funcDecl   []string
	funcCalls  []string
	selectExpr []string
	arrayType  []nodePos
	Fset       *token.FileSet
}

type pkgVisitor struct {
	Fset *token.FileSet
}

type impVisitor struct {
	Fset            *token.FileSet
	relativeImports []string
}

// Get the position of the node in the file
type locate interface {
	getPos(ast.Node) string
}

func (i *impVisitor) getPos(n ast.Node) string {
	return i.Fset.Position(n.Pos()).String()
}

func (fv *fileVisitor) getPos(n ast.Node) string {
	return fv.Fset.Position(n.Pos()).String()
}

func (p *pkgVisitor) getPos(n ast.Node) string {
	return p.Fset.Position(n.Pos()).String()
}

func (c *callVisitor) getPos(n ast.Node) string {
	return c.Fset.Position(n.Pos()).String()
}

type illegal struct {
	T    string
	Name string
	Pos  string
}

func (i *illegal) String() string {
	return i.T + " " + i.Name + " " + i.Pos
}

func getPkgFunc(path string, fsetPkg *token.FileSet) {
	i := &impVisitor{Fset: fsetPkg}
	p := &pkgVisitor{Fset: fsetPkg}
	pkgs, err := parser.ParseDir(fsetPkg, path, nil, parser.AllErrors)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	for pkgname := range pkgs {
		pkg, _ := ast.NewPackage(fsetPkg, pkgs[pkgname].Files, nil, nil)
		pkgName = append(pkgName, pkgname)

		ast.Walk(i, pkg)
		ast.Walk(p, pkg)

		for _, v := range i.relativeImports {
			if isIn(v, openImports) {
				break
			}
			openImports = append(openImports, v)
			getPkgFunc(path+"/"+v, fsetPkg)
		}
	}

}

// reformat from the data base
func splitArgs(args string) []string {
	result := strings.Split(args, " ")
	return result
}

func rightFile(args string) string {
	expectedFiles := splitArgs(args)

	for _, s := range expectedFiles {
		if strings.Contains(s, ".go") {
			return s
		}
	}
	return ""
}

func allowCastingAndImp(allowedImports []string) {
	casted := false
	for i, v := range allowedImports {
		casted = allow(v, casted)
		if v == "--no-array" {
			allArrayTypes = false
			notAllowedArrayT = append(notAllowedArrayT, allowedImports[i+1:]...)
			break
		}
	}
}

type flags struct {
	l struct { // flag for char or string literal
		noLit   bool   // true -> unallows
		pattern string // this pattern
	}
}

// TODO: treat all the flags in this function
// For now, only --no-lit="{PATTERN}"
func parseFlags(args []string) *flags {
	f := &flags{}
	for _, v := range args {
		var flag []string
		if strings.Contains(v, "=") {
			flag = strings.Split(v, "=")
		}
		if flag == nil {
			continue
		}
		if flag[0] == "--no-lit" {
			f.l.noLit = true
			f.l.pattern = flag[1]
		}
	}
	return f
}

func removeAmount(s string) string {
	strRm := strings.TrimFunc(s, func(c rune) bool {
		return c >= '0' && c <= '9' || c == '#'
	})
	return strRm
}

// compares if the function is used a certain amount of times allowed
func allowedAmount(occurrences map[string]int, allowedImports []string) {
	function := ""
	funcSelector := ""
	for _, v := range allowedImports {
		// pkg in case it's a build in function and slice in case it's a selector function
		pkg, slice := trimRelativeImport(v)
		if slice != nil {
			function = strings.Join(slice, ".")
			funcSelector = removeAmount(function)
		} else {
			function = pkg
			funcSelector = removeAmount(pkg)
		}
		if strings.ContainsAny(function, "#") {
			strNbr := strings.TrimPrefix(function, funcSelector+"#")
			nbr, err := strconv.Atoi(strNbr)
			if err != nil {
				log.Panic(err)
			}
			if occurrences[funcSelector] > nbr {
				illegals = append(illegals, illegal{
					T:    "illegal-amount",
					Name: funcSelector + " allowed count " + strNbr + " your count " + strconv.Itoa(occurrences[funcSelector]),
				})
			}
		}
	}
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
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No file or directory")
		return
	}

	var allowedImports []string

	if len(os.Args) > 2 {
		allowedImports = splitArgs(os.Args[2])
	}

	allowCastingAndImp(allowedImports)
	flag := parseFlags(allowedImports)

	filename := strings.TrimSpace(rightFile(os.Args[1]))
	split := strings.Split(filename, "/")
	path := strings.Join(split[:len(split)-1], "/")

	if path == "" {
		path = "."
	}

	fsetFile := token.NewFileSet()
	fsetPkg := token.NewFileSet()

	file, err := parser.ParseFile(fsetFile, filename, nil, parser.AllErrors)

	if err != nil {
		fmt.Println("Parsing")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// Get all the name of all functions declared in the file
	w := &fileVisitor{Fset: fsetFile}
	ast.Walk(w, file)
	getPkgFunc(path, fsetPkg)

	for _, v := range w.funcDecl {
		isFuncAllowed(v, fsetPkg)
	}

	// TODO: Parsing the arguments for the --max-occurrences flag
	allowedAmount(funcOccurrences, allowedImports)

	if flag != nil {
		flag.unallowLits()
	}

	analyzeArrayT()
	analyzeForStmt(allowedImports)

	if len(illegals) > 0 {
		fmt.Println("Cheating")
		for _, i := range illegals {
			fmt.Println(identation + i.String())
		}
		os.Exit(1)
	}
}

func (f flags) unallowLits() {
	if f.l.noLit {
		for _, v := range basicLits {
			if !f.isLitAllowed(v.name) {
				illegals = append(illegals, illegal{
					T:    "illegal-literal",
					Name: v.name,
					Pos:  v.position,
				})
			}
		}
	}
}

func (f flags) isLitAllowed(s string) bool {
	matched, err := regexp.Match(f.l.pattern, []byte(s))

	if err != nil {
		return true
	}

	return !matched
}

func analyzeForStmt(args []string) {
	if isIn("--no-for", args) {
		for _, v := range forStmts {
			illegals = append(illegals, illegal{
				T:    "illegal-loop",
				Name: v.name,
				Pos:  v.position,
			})
		}
	}
}

func analyzeArrayT() {
	if !allArrayTypes {
		l := len(notAllowedArrayT)
		for _, v := range arraysInstances {
			if l == 0 ||
				isIn(v.name, notAllowedArrayT) {
				illegals = append(illegals, illegal{
					T:    "illegal-array-type",
					Name: v.name,
					Pos:  v.position,
				})
			}
		}
	}
}

func allow(s string, casted bool) bool {
	if strings.ContainsRune(s, '.') {
		allowImport(s)
	} else {
		if allowedFun == nil {
			allowedFun = make(map[string]bool)
		}
		if strings.ContainsAny(s, "#") {
			s = removeAmount(s)
		}
		allowedFun[s] = true
	}

	if s == "--cast" && !casted {
		for _, v := range predeclaredTypes {
			allow(v, false)
		}
		return true
	}
	return casted
}

// Returns true if the string matches the format of a relative import
func isRelativeImport(s string) bool {
	relativeImport, _ := regexp.MatchString(`\.\.\\??`, s)
	return relativeImport
}

// Returns true if the string represents an import package
// i.e and expresion like <lib>.<function>
func isImport(s string) bool {
	matched, _ := regexp.MatchString(`.\..`, s)
	return matched
}

func trimRelativeImport(str string) (string, []string) {
	var pkg string
	var slice []string
	if isImport(str) && !isRelativeImport(str) {
		splited := strings.Split(str, "/")
		slice = strings.Split(splited[len(splited)-1], ".")
		splited[len(splited)-1] = slice[0]
		pkg = strings.Join(splited, "/")

		if allowedImp[slice[0]] == nil {
			allowedImp[slice[0]] = make(map[string]bool)
		}
		fn := slice[len(slice)-1]
		allowedImp[slice[0]][fn] = true
	} else {
		pkg = str
	}
	return pkg, slice
}

func allowImport(s string) {
	if allowedImp == nil {
		allowedImp = make(map[string]map[string]bool)
	}
	if allImports == nil {
		allImports = make(map[string]bool)
	}

	pkg, slice := trimRelativeImport(s)

	allImports[pkg] = true

	fn := "*"
	if len(slice) > 1 {
		fn = removeAmount(slice[1])
	}

	if allowedImp[pkg] == nil {
		allowedImp[pkg] = make(map[string]bool)
	}
	allowedImp[pkg][fn] = true
}

func addToIllegals(funcname string) {
	for _, v := range callX {
		if v.name == funcname {
			pos := v.position
			setIllegal("illegal-call", v.name, pos)
		}
	}
	if funcDeclPkg[funcname] != nil {
		pos := funcDeclPkg[funcname].position
		setIllegal("illegal-call", funcname, pos)

	}
}

// First ignoring the imported functions
func isFuncAllowed(funcname string, fset *token.FileSet) bool {
	if allowedFun[funcname] {
		return true
	}

	if !isFuncDeclIn(funcname, funcDeclPkg) {
		addToIllegals(funcname)
		return false
	}
	bodyf := funcDeclPkg[funcname].body

	if bodyf == nil {
		addToIllegals(funcname)
		fmt.Println("Body is nil")
		return false
	}
	c := &callVisitor{Fset: fset}
	ast.Walk(c, bodyf)

	res := true

	for _, v := range c.Calls {
		if v == funcname {
			continue
		}
		allowed := isFuncAllowed(v, fset)
		if !allowed {
			addToIllegals(funcname)
			for _, v := range c.Calls {
				if v == funcname {
					continue
				}
				allowed := isFuncAllowed(v, fset)
				if !allowed {
					addToIllegals(funcname)
				}
				res = res && allowed
			}
		}
		res = res && allowed
	}
	return res
}

func isFuncDeclIn(funcname string, fundecl map[string]*funcBody) bool {
	return fundecl[funcname] != nil
}

func isFuncCallIn(funcname string, funP []nodePos) bool {
	for _, v := range funP {
		if v.name == funcname {
			return true
		}
	}
	return false
}

func isIn(s string, slc []string) bool {
	for _, v := range slc {
		if s == v {
			return true
		}
	}
	return false
}

// Keeps the positions of each function call and declaration
type funcBody struct {
	position string
	body     *ast.BlockStmt
}

type nodePos struct {
	position string
	name     string
}

func (fv *fileVisitor) Visit(n ast.Node) ast.Visitor {
	fv.checkImport(n)
	if ide, ok := n.(*ast.CallExpr); ok {
		if opt, ok := ide.Fun.(*ast.Ident); ok {
			fv.funcCalls = append(fv.funcCalls, opt.Name)
			newFun := nodePos{position: fv.getPos(n), name: opt.Name}
			callX = append(callX, newFun)
		}
	}

	if expr, ok := n.(*ast.SelectorExpr); ok {
		if x, ok := expr.X.(*ast.Ident); ok {
			fv.selectExpr = append(fv.selectExpr, x.Name+"."+expr.Sel.Name)

			// saves the function in to the map, from the package
			if importPkg[x.Name] != nil {
				importPkg[x.Name].functions = append(importPkg[x.Name].functions, x.Name+"."+expr.Sel.Name)
			}
			if funcOccurrences == nil {
				funcOccurrences = make(map[string]int)
			}
			funcOccurrences[x.Name+"."+expr.Sel.Name]++
		}
	}

	if ex, ok := n.(*ast.ArrayType); ok {
		if op, ok := ex.Elt.(*ast.Ident); ok {
			fv.arrayType = append(fv.arrayType, nodePos{
				name:     op.Name,
				position: fv.getPos(n),
			})
		}
	}

	if exp, ok := n.(*ast.FuncDecl); ok {
		fv.funcDecl = append(fv.funcDecl, exp.Name.Name)
		for _, v := range exp.Type.Params.List {
			if _, ok := v.Type.(*ast.FuncType); ok {
				if allowedFun == nil {
					allowedFun = make(map[string]bool)
				}
				for _, name := range v.Names {
					allowedFun[name.Name] = true
				}
			}
		}
	}

	if ex, ok := n.(*ast.AssignStmt); ok {
		if exp, ok := ex.Rhs[0].(*ast.FuncLit); ok {
			if ide, ok := ex.Lhs[0].(*ast.Ident); ok {
				if funcDeclPkg == nil {
					funcDeclPkg = make(map[string]*funcBody)
				}
				funcDeclPkg[ide.Name] = &funcBody{body: exp.Body, position: fv.getPos(n)}
			}
		}
	}

	return fv
}

func positionIsIn(illegals []illegal, pos string) bool {
	for _, v := range illegals {
		if v.Pos == pos {
			return true
		}
	}
	return false
}

func (p *pkgVisitor) Visit(n ast.Node) ast.Visitor {
	if exp, ok := n.(*ast.FuncDecl); ok {
		if funcDeclPkg == nil {
			funcDeclPkg = make(map[string]*funcBody)
		}
		funcDeclPkg[exp.Name.Name] = &funcBody{body: exp.Body, position: p.getPos(n)}
		for _, pkg := range pkgName {
			if importPkg[pkg] != nil && isIn(pkg+"."+exp.Name.Name, importPkg[pkg].functions) {
				funcDeclPkg[pkg+"."+exp.Name.Name] = &funcBody{body: exp.Body, position: p.getPos(n)}
			}
		}
	}

	if ex, ok := n.(*ast.AssignStmt); ok {
		if exp, ok := ex.Rhs[0].(*ast.FuncLit); ok {
			if ide, ok := ex.Lhs[0].(*ast.Ident); ok {
				if funcDeclPkg == nil {
					funcDeclPkg = make(map[string]*funcBody)
				}
				funcDeclPkg[ide.Name] = &funcBody{body: exp.Body, position: p.getPos(n)}
			}
		}
	}
	return p
}

func setIllegal(illegalType, funcName, pos string) {
	if !positionIsIn(illegals, pos) {
		illegals = append(illegals, illegal{
			T:    illegalType,
			Name: funcName,
			Pos:  pos,
		})
	}
}

// Signals that exists at least one callExpr in the node
func (c *callVisitor) Visit(n ast.Node) ast.Visitor {
	if id, ok := n.(*ast.BasicLit); ok {
		if id.Kind != token.CHAR && id.Kind != token.STRING {
			return nil
		}
		basicLits = append(basicLits, nodePos{position: c.getPos(n), name: id.Value})
	}

	if exp, ok := n.(*ast.CallExpr); ok {
		if fun, ok := exp.Fun.(*ast.Ident); ok {
			c.Calls = append(c.Calls, fun.Name)
			newFun := nodePos{position: c.getPos(n), name: fun.Name}
			callX = append(callX, newFun)
			if funcOccurrences == nil {
				funcOccurrences = make(map[string]int)
			}
			funcOccurrences[fun.Name]++
			return c
		}
	}
	// SelectorExpr is when we access a value (dot opperator)
	// We need to check those for specific functions
	if expr, ok := n.(*ast.SelectorExpr); ok {
		x, ok := expr.X.(*ast.Ident)
		if !ok {
			// in this case we are deep in an access
			// example, fmt is banned, but pouet isn't.
			// we must allow pouet.fmt.x but not fmt.x
			// this is the pouet.fmt.x case.
			return c
		}
		pkg := allowedImp[x.Name]
		f := x.Name + "." + expr.Sel.Name
		if funcDeclPkg[f] != nil {
			c.Calls = append(c.Calls, f)
			return c
		}

		if pkg == nil {
			if allImports[x.Name] {
				pos := c.getPos(n)
				setIllegal("illegal-access", f, pos)
			}
			return c
		}
		if !pkg["*"] && !pkg[expr.Sel.Name] {
			// all the package is not whiteList and is not explicitly allowed
			pos := c.getPos(n)
			setIllegal("illegal-access", f, pos)
		}
	}

	if ex, ok := n.(*ast.ArrayType); ok {
		if op, ok := ex.Elt.(*ast.Ident); ok {
			arraysInstances = append(arraysInstances, nodePos{
				name:     op.Name,
				position: c.getPos(n),
			})
		}
	}

	if _, ok := n.(*ast.ForStmt); ok {
		forStmts = append(forStmts, nodePos{
			name:     "for",
			position: c.getPos(n),
		})
	}

	return c
}

func (fv *fileVisitor) checkImport(n ast.Node) ast.Visitor {
	if spec, ok := n.(*ast.ImportSpec); ok {
		pkg := spec.Path.Value[1 : len(spec.Path.Value)-1]
		if allowedImp[pkg] == nil {
			pos := fv.getPos(n)
			setIllegal("illegal-import", pkg, pos)
			return fv
		}
		// if the import is named, we need to move it to the new name
		name := ""
		if spec.Name != nil {
			name = spec.Name.Name
		} else if strings.ContainsRune(pkg, '/') {
			parts := strings.Split(pkg, "/")
			name = parts[len(parts)-1]
		}
		if allowedImp[pkg] != nil {
			if name != "" {
				allowedImp[name] = allowedImp[pkg]
				allowedImp[pkg] = nil
			}
		}

		if isRelativeImport(pkg) {
			if importPkg == nil {
				importPkg = make(map[string]*pkgFunc)
			}
			if name != "" {
				importPkg[name] = &pkgFunc{
					path: pkg,
				}
			}
			relativeImports = append(relativeImports, name)
		}
	}
	return fv
}

func (i *impVisitor) Visit(n ast.Node) ast.Visitor {
	if spec, ok := n.(*ast.ImportSpec); ok {
		pkg := spec.Path.Value[1 : len(spec.Path.Value)-1]
		if allImports == nil {
			allImports = make(map[string]bool)
		}
		pkgSplit := strings.Split(pkg, "/")

		allImports[pkg] = true
		allImports[pkgSplit[len(pkgSplit)-1]] = true

		if isRelativeImport(pkg) {
			i.relativeImports = append(i.relativeImports, pkg)
		}
	}
	return i
}
