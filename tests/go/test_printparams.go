package main

import (
	"strings"

	"github.com/01-edu/z01"
)

func main() {
	table := append(z01.MultRandWords(), "choumi is the best cat")
	for _, s := range table {
		z01.ChallengeMain(strings.Fields(s)...)
	}
}
