package main

import (
	"strings"

	"../lib"
)

func main() {
	table := append(lib.MultRandWords(), " ")
	table = append(table, "1")
	table = append(table, "1 2")

	for _, s := range table {
		lib.ChallengeMain("displaylastparam", strings.Fields(s)...)
	}
}
