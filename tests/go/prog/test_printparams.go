package main

import (
	"strings"

	"../lib"
)

func main() {
	table := append(lib.MultRandWords(), "choumi is the best cat")
	for _, s := range table {
		lib.ChallengeMain("printparams", strings.Fields(s)...)
	}
}
