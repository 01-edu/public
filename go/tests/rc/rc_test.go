package main

import (
	"sort"
	"strings"
	"testing"

	"github.com/01-edu/z01"
)

func TestFlags(t *testing.T) {
	argsAndSolution := map[string]string{
		"tests/testingSimpleFunc.go": `Cheating:
	TYPE:             	NAME:             	LOCATION:
	illegal-import    	regexp            	tests/testingSimpleFunc.go:4:2
	illegal-call      	len               	tests/testingSimpleFunc.go:10:9
	illegal-access    	regexp.MustCompile	tests/testingSimpleFunc.go:8:8
	illegal-definition	SimpleFunc        	tests/testingSimpleFunc.go:7:1
`,

		"-no-for -no-lit=[a-z] tests/printalphabet/printalphabet.go": `Cheating:
	TYPE:             	NAME:                	LOCATION:
	illegal-import    	fmt                  	tests/printalphabet/printalphabet.go:4:2
	illegal-import    	github.com/01-edu/z01	tests/printalphabet/printalphabet.go:6:2
	illegal-call      	append               	tests/printalphabet/printalphabet.go:11:7
	illegal-definition	fillArray            	tests/printalphabet/printalphabet.go:9:1
	illegal-call      	int                  	tests/printalphabet/printalphabet.go:17:7
	illegal-access    	z01.PrintRune        	tests/printalphabet/printalphabet.go:19:3
	illegal-access    	z01.PrintRune        	tests/printalphabet/printalphabet.go:21:2
	illegal-definition	main                 	tests/printalphabet/printalphabet.go:15:1
	illegal-access    	fmt.Println          	tests/printalphabet/printalphabet.go:26:3
	illegal-definition	defFun               	tests/printalphabet/printalphabet.go:25:2
	illegal-call      	defFun               	tests/printalphabet/printalphabet.go:28:2
	illegal-definition	testingScope         	tests/printalphabet/printalphabet.go:24:1
	illegal-loop      	for                  	tests/printalphabet/printalphabet.go:10:2
	illegal-lit       	'a'                  	tests/printalphabet/printalphabet.go:10:11
	illegal-lit       	'z'                  	tests/printalphabet/printalphabet.go:10:21
	illegal-lit       	'a'                  	tests/printalphabet/printalphabet.go:16:14
	illegal-lit       	'b'                  	tests/printalphabet/printalphabet.go:16:19
	illegal-lit       	'c'                  	tests/printalphabet/printalphabet.go:16:24
	illegal-lit       	'd'                  	tests/printalphabet/printalphabet.go:16:29
	illegal-lit       	'e'                  	tests/printalphabet/printalphabet.go:16:34
	illegal-lit       	'f'                  	tests/printalphabet/printalphabet.go:16:39
	illegal-lit       	'a'                  	tests/printalphabet/printalphabet.go:17:11
	illegal-lit       	'\n'                 	tests/printalphabet/printalphabet.go:21:16
	illegal-lit       	"Hello"              	tests/printalphabet/printalphabet.go:28:9
`,
		"-cast tests/eightqueens.go": `Cheating:
	TYPE:             	NAME:                	LOCATION:
	illegal-import    	github.com/01-edu/z01	tests/eightqueens.go:4:2
	illegal-access    	z01.PrintRune        	tests/eightqueens.go:49:5
	illegal-access    	z01.PrintRune        	tests/eightqueens.go:55:2
	illegal-definition	printQueens          	tests/eightqueens.go:42:1
`,
		"-no-arrays tests/printalphabet/printalphabet.go": `Cheating:
	TYPE:             	NAME:                	LOCATION:
	illegal-import    	fmt                  	tests/printalphabet/printalphabet.go:4:2
	illegal-import    	github.com/01-edu/z01	tests/printalphabet/printalphabet.go:6:2
	illegal-call      	append               	tests/printalphabet/printalphabet.go:11:7
	illegal-definition	fillArray            	tests/printalphabet/printalphabet.go:9:1
	illegal-call      	int                  	tests/printalphabet/printalphabet.go:17:7
	illegal-access    	z01.PrintRune        	tests/printalphabet/printalphabet.go:19:3
	illegal-access    	z01.PrintRune        	tests/printalphabet/printalphabet.go:21:2
	illegal-definition	main                 	tests/printalphabet/printalphabet.go:15:1
	illegal-access    	fmt.Println          	tests/printalphabet/printalphabet.go:26:3
	illegal-definition	defFun               	tests/printalphabet/printalphabet.go:25:2
	illegal-call      	defFun               	tests/printalphabet/printalphabet.go:28:2
	illegal-definition	testingScope         	tests/printalphabet/printalphabet.go:24:1
	illegal-slice     	rune                 	tests/printalphabet/printalphabet.go:9:18
	illegal-slice     	rune                 	tests/printalphabet/printalphabet.go:16:7
`,
		"-no-slices tests/printalphabet/printalphabet.go": `Cheating:
	TYPE:             	NAME:                	LOCATION:
	illegal-import    	fmt                  	tests/printalphabet/printalphabet.go:4:2
	illegal-import    	github.com/01-edu/z01	tests/printalphabet/printalphabet.go:6:2
	illegal-call      	append               	tests/printalphabet/printalphabet.go:11:7
	illegal-definition	fillArray            	tests/printalphabet/printalphabet.go:9:1
	illegal-call      	int                  	tests/printalphabet/printalphabet.go:17:7
	illegal-access    	z01.PrintRune        	tests/printalphabet/printalphabet.go:19:3
	illegal-access    	z01.PrintRune        	tests/printalphabet/printalphabet.go:21:2
	illegal-definition	main                 	tests/printalphabet/printalphabet.go:15:1
	illegal-access    	fmt.Println          	tests/printalphabet/printalphabet.go:26:3
	illegal-definition	defFun               	tests/printalphabet/printalphabet.go:25:2
	illegal-call      	defFun               	tests/printalphabet/printalphabet.go:28:2
	illegal-definition	testingScope         	tests/printalphabet/printalphabet.go:24:1
	illegal-slice     	rune                 	tests/printalphabet/printalphabet.go:9:18
	illegal-slice     	rune                 	tests/printalphabet/printalphabet.go:16:7
`,
		"-no-these-slices=int,rune tests/printalphabet/printalphabet.go": `Cheating:
	TYPE:             	NAME:                	LOCATION:
	illegal-import    	fmt                  	tests/printalphabet/printalphabet.go:4:2
	illegal-import    	github.com/01-edu/z01	tests/printalphabet/printalphabet.go:6:2
	illegal-call      	append               	tests/printalphabet/printalphabet.go:11:7
	illegal-definition	fillArray            	tests/printalphabet/printalphabet.go:9:1
	illegal-call      	int                  	tests/printalphabet/printalphabet.go:17:7
	illegal-access    	z01.PrintRune        	tests/printalphabet/printalphabet.go:19:3
	illegal-access    	z01.PrintRune        	tests/printalphabet/printalphabet.go:21:2
	illegal-definition	main                 	tests/printalphabet/printalphabet.go:15:1
	illegal-access    	fmt.Println          	tests/printalphabet/printalphabet.go:26:3
	illegal-definition	defFun               	tests/printalphabet/printalphabet.go:25:2
	illegal-call      	defFun               	tests/printalphabet/printalphabet.go:28:2
	illegal-definition	testingScope         	tests/printalphabet/printalphabet.go:24:1
	illegal-slice     	rune                 	tests/printalphabet/printalphabet.go:9:18
	illegal-slice     	rune                 	tests/printalphabet/printalphabet.go:16:7
`,
		`-allow-builtin tests/doopprog/main.go fmt.Println strconv.Atoi os.Args`: ``,
		`-cast tests/doopprog/main.go fmt.Println strconv.Atoi os.Args len`:      ``,
		`tests/testingWrapping.go`: `Cheating:
	TYPE:             	NAME:           	LOCATION:
	illegal-call      	len             	tests/utilDepth2/wrapper.go:4:9
	illegal-definition	LenWrapper      	tests/utilDepth2/wrapper.go:3:1
	illegal-access    	util2.LenWrapper	tests/util/util.go:10:9
	illegal-definition	LenWrapperU     	tests/util/util.go:9:1
	illegal-access    	util.LenWrapperU	tests/testingWrapping.go:8:9
	illegal-definition	Length          	tests/testingWrapping.go:7:1
`,
		`tests/testingWrapping.go len`: ``,
		`tests/empty/empty len`: `	No file to analyze
`,
		`tests/empty/empty.go tests/empty/empty`: `	tests/empty/empty.go:1:1: expected ';', found 'EOF' (and 2 more errors)
`,
	}
	Compare(t, argsAndSolution)
}

func Compare(t *testing.T, argsAndSol map[string]string) {
	for args, sol := range argsAndSol {
		a := strings.Split(args, " ")
		out, err := z01.MainOut("../rc", a...)
		if !EqualResult(sol, out) {
			if err == nil {
				t.Errorf("./rc %s prints %q\n instead of %q\n", args, out, sol)
			}
			if err != nil && !EqualResult(sol, err.Error()) {
				t.Errorf("./rc %s prints %q\n instead of %q\n", args, err.Error(), sol)
			}
		}
	}
}

func EqualResult(sol, out string) bool {
	// split
	solSli := strings.Split(sol, "\n")
	outSli := strings.Split(out, "\n")
	// sort
	sort.Sort(sort.StringSlice(solSli))
	sort.Sort(sort.StringSlice(outSli))
	// join
	sol = strings.Join(solSli, "\n")
	out = strings.Join(outSli, "\n")
	// compare
	return sol == out
}
