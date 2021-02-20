package main

import (
	"strings"

	"github.com/01-edu/public/go/tests/lib"
)

func main() {
	table := append(lib.MultRandWords(),
		" ",
		"1",
		"1 2",
		"1 2 3",
	)
	for _, s := range table {
		lib.ChallengeMain("displaylastparam", strings.Fields(s)...)
	}
}
