package main

import (
	"strings"

	"lib"
)

func main() {
	table := append(lib.MultRandWords(),
		" ",
		"1",
		"1 2",
		"1 2 3",
	)
	for _, s := range table {
		lib.ChallengeMain("displayfirstparam", strings.Fields(s)...)
	}
}
