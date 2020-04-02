package main

import (
	"testing"

	"github.com/01-edu/z01"
)

func TestSearchReplace(t *testing.T) {

	type nodeTest struct {
		dataSearched    string
		letterLookedFor string
		letterReplacing string
	}

	table := []nodeTest{}

	for i := 0; i < 20; i++ {

		letter1 := []rune(z01.RandAlnum())
		letter2 := []rune(z01.RandAlnum())

		table = append(table,
			nodeTest{
				dataSearched:    z01.RandWords(),
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
		z01.ChallengeMainExam(t, arg.dataSearched, arg.letterLookedFor, arg.letterReplacing)
	}
}
