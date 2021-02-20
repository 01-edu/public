package main

import (
	"strings"

	"github.com/01-edu/public/go/tests/lib"
)

func main() {
	table := append(lib.MultRandWords(), "1 2 3 4 5")

	for _, s := range table {
		lib.ChallengeMain("boolean", strings.Fields(s)...)
	}
}
