package main

import "github.com/01-edu/z01"

func main() {
	table := append(z01.MultRandWords(),
		"",
		"             a as",
		"   f     d",
		"   asd    ad",
		"   salut !!! ",
		" salut ! ! !",
		"salut ! !",
	)
	for _, s := range table {
		z01.ChallengeMain("firstword", s)
	}
}
