package main

import (
	"strings"

	"github.com/01-edu/z01"
)

func main() {
	table := append(z01.MultRandWords(), "1 2 3 4 5")

	for _, s := range table {
		z01.ChallengeMain(strings.Fields(s)...)
	}
}
