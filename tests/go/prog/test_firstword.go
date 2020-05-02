package main

import "../lib"

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
