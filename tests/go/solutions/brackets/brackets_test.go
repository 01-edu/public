package main

import (
	"testing"

	"github.com/01-edu/z01"
)

func TestBrackets(t *testing.T) {
	oneArgs := []string{"(johndoe)", ")()", "([)]", "{2*[d - 3]/(12)}"}

	// 18 random tests ( at least half are valid)
	for i := 0; i < 3; i++ {
		oneArgs = append(oneArgs, "("+z01.RandASCII()+")")
		oneArgs = append(oneArgs, "["+z01.RandASCII()+"]")
		oneArgs = append(oneArgs, "{"+z01.RandASCII()+"}")
		oneArgs = append(oneArgs, "("+z01.RandAlnum()+")")
		oneArgs = append(oneArgs, "["+z01.RandAlnum()+"]")
		oneArgs = append(oneArgs, "{"+z01.RandAlnum()+"}")
	}

	// No args testig
	z01.ChallengeMainExam(t)

	for _, v := range oneArgs {
		z01.ChallengeMainExam(t, v)
	}

	arg1 := []string{"", "{[(0 + 0)(1 + 1)](3*(-1)){()}}"}
	arg2 := []string{"{][]}", "{3*[21/(12+ 23)]}"}
	arg3 := []string{"{([)])}", "{{{something }- [something]}}", "there are"}
	multArg := [][]string{arg1, arg2, arg3}

	for _, v := range multArg {
		z01.ChallengeMainExam(t, v...)
	}
}
