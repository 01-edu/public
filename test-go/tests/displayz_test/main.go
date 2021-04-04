package main

import (
	"strings"

	"github.com/01-edu/public/test-go/lib"
)

func main() {
	table := append(lib.MultRandWords(),
		"dsfdz",
		"",
		"1",
		"1",
	)
	for _, s := range table {
		lib.ChallengeMain("displayz", strings.Fields(s)...)
	}

	lib.ChallengeMain("displayz", "1", "z")
}
