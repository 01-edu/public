package main

import (
	"strings"

	"github.com/01-edu/public/go/tests/lib"
)

func main() {
	table := append(lib.MultRandWords(),
		"dsfda",
		"",
		"1",
		"1",
	)
	for _, s := range table {
		lib.ChallengeMain("displaya", strings.Fields(s)...)
	}
	lib.ChallengeMain("displaya", "1", "a")
}
