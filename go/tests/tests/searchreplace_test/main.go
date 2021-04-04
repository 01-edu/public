package main

import (
	"github.com/01-edu/public/go/tests/lib"
)

func main() {
	type nodeTest struct {
		dataSearched    string
		letterLookedFor string
		letterReplacing string
	}

	table := []nodeTest{}

	for i := 0; i < 20; i++ {
		letter1 := []rune(lib.RandAlnum())
		letter2 := []rune(lib.RandAlnum())

		table = append(table,
			nodeTest{
				dataSearched:    lib.RandWords(),
				letterLookedFor: string(letter1[0]),
				letterReplacing: string(letter2[0]),
			})
	}

	table = append(table,
		nodeTest{
			dataSearched:    "hélla",
			letterLookedFor: "é",
			letterReplacing: "o",
		},
		nodeTest{
			dataSearched:    "hella",
			letterLookedFor: "z",
			letterReplacing: "o",
		},
		nodeTest{
			dataSearched:    "hella",
			letterLookedFor: "h",
			letterReplacing: "o",
		},
	)

	for _, arg := range table {
		lib.ChallengeMain("searchreplace", arg.dataSearched, arg.letterLookedFor, arg.letterReplacing)
	}
}
