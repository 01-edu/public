package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/01-edu/z01"
)

func challengeProgram(t *testing.T, stuPath, solPath string, args ...string) {
	exercise := strings.ToLower(
		strings.TrimPrefix(t.Name(), "Test"))
	out, err := z01.MainOut(stuPath+exercise, args...)
	if err != nil {
		t.Error(err)
	}
	correct, err := z01.MainOut(solPath+exercise, args...)
	if err != nil {
		t.Error(err)
	}
	if out != correct {
		t.Errorf("./%s %s prints %q instead of %q\n",
			exercise, strings.Join(args, " "), out, correct)
	}
}

func TestFlags(t *testing.T) {
	argsAndSolution := map[string]string{
		"tests/alphacount.go": `Parsing
  OK
Cheating
  illegal-import regexp tests/alphacount.go:4:2
  illegal-access regexp.MustCompile tests/alphacount.go:8:8
  illegal-call len tests/alphacount.go:10:9
  illegal-call AlphaCount tests/alphacount.go:7:1
`,

		"-no-for -no-lit=[a-z] tests/printalphabet/printalphabet.go": `Parsing
  OK
Cheating
  illegal-import github.com/01-edu/z01 tests/printalphabet/printalphabet.go:4:2
  illegal-call append tests/printalphabet/printalphabet.go:9:7
  illegal-call fillArray tests/printalphabet/printalphabet.go:15:2
  illegal-call fillArray tests/printalphabet/printalphabet.go:7:1
  illegal-access z01.PrintRune tests/printalphabet/printalphabet.go:17:3
  illegal-access z01.PrintRune tests/printalphabet/printalphabet.go:19:2
  illegal-call main tests/printalphabet/printalphabet.go:13:1
  illegal-literal 'a' tests/printalphabet/printalphabet.go:8:11
  illegal-literal 'z' tests/printalphabet/printalphabet.go:8:21
  illegal-literal '\n' tests/printalphabet/printalphabet.go:19:16
  illegal-literal 'a' tests/printalphabet/printalphabet.go:8:11
  illegal-literal 'z' tests/printalphabet/printalphabet.go:8:21
  illegal-literal 'a' tests/printalphabet/printalphabet.go:8:11
  illegal-literal 'z' tests/printalphabet/printalphabet.go:8:21
  illegal-loop for tests/printalphabet/printalphabet.go:8:2
  illegal-loop for tests/printalphabet/printalphabet.go:8:2
  illegal-loop for tests/printalphabet/printalphabet.go:8:2
`,
		"tests/eightqueens.go": `Parsing
  OK
Cheating
  illegal-import github.com/01-edu/z01 tests/eightqueens.go:4:2
  illegal-access z01.PrintRune tests/eightqueens.go:49:5
  illegal-access z01.PrintRune tests/eightqueens.go:55:2
  illegal-call rune tests/eightqueens.go:49:19
  illegal-call printQueens tests/eightqueens.go:70:5
  illegal-call printQueens tests/eightqueens.go:42:1
  illegal-call tryX tests/eightqueens.go:73:5
  illegal-call tryX tests/eightqueens.go:85:2
  illegal-call tryX tests/eightqueens.go:60:1
  illegal-call EightQueens tests/eightqueens.go:83:1
`,
		"-cast tests/eightqueens.go": `Parsing
  OK
Cheating
  illegal-import github.com/01-edu/z01 tests/eightqueens.go:4:2
  illegal-access z01.PrintRune tests/eightqueens.go:49:5
  illegal-access z01.PrintRune tests/eightqueens.go:55:2
`,
		"-no-arrays tests/printalphabet/printalphabet.go ": `Parsing
  OK
Cheating
  illegal-import github.com/01-edu/z01 tests/printalphabet/printalphabet.go:4:2
  illegal-call append tests/printalphabet/printalphabet.go:9:7
  illegal-call fillArray tests/printalphabet/printalphabet.go:15:2
  illegal-call fillArray tests/printalphabet/printalphabet.go:7:1
  illegal-access z01.PrintRune tests/printalphabet/printalphabet.go:17:3
  illegal-access z01.PrintRune tests/printalphabet/printalphabet.go:19:2
  illegal-call main tests/printalphabet/printalphabet.go:13:1
  illegal-array-type rune tests/printalphabet/printalphabet.go:14:8
`,
		"-no-these-arrays rune tests/printalphabet/printalphabet.go ": `Parsing
  OK
Cheating
  illegal-import github.com/01-edu/z01 tests/printalphabet/printalphabet.go:4:2
  illegal-call append tests/printalphabet/printalphabet.go:9:7
  illegal-call fillArray tests/printalphabet/printalphabet.go:15:2
  illegal-call fillArray tests/printalphabet/printalphabet.go:7:1
  illegal-access z01.PrintRune tests/printalphabet/printalphabet.go:17:3
  illegal-access z01.PrintRune tests/printalphabet/printalphabet.go:19:2
  illegal-call main tests/printalphabet/printalphabet.go:13:1
  illegal-array-type rune tests/printalphabet/printalphabet.go:14:8
`,
		`-no-these-arrays "int" tests/printalphabet/printalphabet.go`: `Parsing
  OK
Cheating
  illegal-import github.com/01-edu/z01 tests/printalphabet/printalphabet.go:4:2
  illegal-call append tests/printalphabet/printalphabet.go:9:7
  illegal-call fillArray tests/printalphabet/printalphabet.go:15:2
  illegal-call fillArray tests/printalphabet/printalphabet.go:7:1
  illegal-access z01.PrintRune tests/printalphabet/printalphabet.go:17:3
  illegal-access z01.PrintRune tests/printalphabet/printalphabet.go:19:2
  illegal-call main tests/printalphabet/printalphabet.go:13:1
`,
	}

	Compare(t, argsAndSolution)
}

func Compare(t *testing.T, argsAndSol map[string]string) {
	for args, sol := range argsAndSol {
		out, err := z01.MainOut("../rc", strings.Split(args, " ")...)
		if err.Error() != sol && out != sol {
			fmt.Println("Error:", err)
			fmt.Println("Out:", out)
			t.Errorf("./rc %s prints %q\n instead of %q\n", args, out, sol)
		}
	}
}

func TestWrapping(t *testing.T) {
	argsAndSolution := map[string]string{
		"tests/testingWrapping.go ./util": `Parsing
  OK
Cheating
  illegal-call len tests/util/util.go:6:9
  illegal-call util.LenWrapper tests/util/util.go:5:1
  illegal-call Length tests/testingWrapping.go:7:1
`,
	}
	Compare(t, argsAndSolution)
}
