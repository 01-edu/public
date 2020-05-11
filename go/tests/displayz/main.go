package main

import (
	"strings"

	"github.com/01-edu/public/go/lib"
)

func main() {
	table := append(lib.MultRandWords(), "dsfdz")
	table = append(table, "")
	table = append(table, "1")
	table = append(table, "1")

	for _, s := range table {
		lib.ChallengeMain("displayz", strings.Fields(s)...)
	}

	lib.ChallengeMain("displayz", "1", "z")
}
