package main

import (
	"strings"

	"github.com/01-edu/public/go/lib"
)

func main() {
	table := append(lib.MultRandWords(), "dsfda")
	table = append(table, "")
	table = append(table, "1")
	table = append(table, "1")

	for _, s := range table {
		lib.ChallengeMain("displaya", strings.Fields(s)...)
	}

	lib.ChallengeMain("displaya", "1", "a")
}
