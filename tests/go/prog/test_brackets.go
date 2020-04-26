package main

import "github.com/01-edu/z01"

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
			"("+z01.RandASCII()+")",
			"["+z01.RandASCII()+"]",
			"{"+z01.RandASCII()+"}",
			"("+z01.RandAlnum()+")",
			"["+z01.RandAlnum()+"]",
			"{"+z01.RandAlnum()+"}",
		)
	}

	z01.ChallengeMain("brackets")

	for _, v := range oneArgs {
		z01.ChallengeMain("brackets", v)
	}

	multArg := [][]string{
		{"", "{[(0 + 0)(1 + 1)](3*(-1)){()}}"},
		{"{][]}", "{3*[21/(12+ 23)]}"},
		{"{([)])}", "{{{something }- [something]}}", "there are"},
	}

	for _, v := range multArg {
		z01.ChallengeMain("brackets", v...)
	}
}
