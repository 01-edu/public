package main

import (
	"github.com/01-edu/public/test-go/lib"
)

func main() {
	oneArgs := []string{
		"(johndoe)",
		")()",
		"([)]",
		"{2*[d - 3]/(12)}",
	}

	// 18 random tests ( at least half are valid)
	for i := 0; i < 3; i++ {
		oneArgs = append(oneArgs,
			"("+lib.RandASCII()+")",
			"["+lib.RandASCII()+"]",
			"{"+lib.RandASCII()+"}",
			"("+lib.RandAlnum()+")",
			"["+lib.RandAlnum()+"]",
			"{"+lib.RandAlnum()+"}",
		)
	}

	lib.ChallengeMain("brackets")

	for _, v := range oneArgs {
		lib.ChallengeMain("brackets", v)
	}

	multArg := [][]string{
		{"", "{[(0 + 0)(1 + 1)](3*(-1)){()}}"},
		{"{][]}", "{3*[21/(12+ 23)]}"},
		{"{([)])}", "{{{something }- [something]}}", "there are"},
	}

	for _, v := range multArg {
		lib.ChallengeMain("brackets", v...)
	}
}
