package main

import (
	"strings"

	"github.com/01-edu/z01"
)

func main() {
	table := append(z01.MultRandWords(), " ")
	table = append(table, "1")
	table = append(table, "1 2")
	table = append(table, "1 2 3")

	for _, s := range table {
		z01.ChallengeMain(strings.Fields(s)...)
	}
}
