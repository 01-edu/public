package main

import (
	"strings"

	"github.com/01-edu/z01"
)

func main() {
	table := append(z01.MultRandWords(), "dsfdz")
	table = append(table, "")
	table = append(table, "1")
	table = append(table, "1")

	for _, s := range table {
		z01.ChallengeMain("displayz", strings.Fields(s)...)
	}

	z01.ChallengeMain("displayz", "1", "z")
}
