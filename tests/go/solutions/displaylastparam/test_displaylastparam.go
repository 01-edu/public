package main

import (
	"strings"

	"github.com/01-edu/z01"
)

func main() {
	table := append(z01.MultRandWords(), " ")
	table = append(table, "1")
	table = append(table, "1 2")

	for _, s := range table {
		z01.ChallengeMain("displaylastparam", strings.Fields(s)...)
	}
}
