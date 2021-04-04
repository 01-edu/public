package main

import (
	"github.com/01-edu/public/test-go/lib"
)

func main() {
	table := append(lib.MultRandWords(),
		"",
		"             a as",
		"   f     d",
		"   asd    ad",
		"   salut !!! ",
		" salut ! ! !",
		"salut ! !",
	)
	for _, s := range table {
		lib.ChallengeMain("firstword", s)
	}
}
