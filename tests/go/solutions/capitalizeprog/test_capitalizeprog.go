package main

import (
	"github.com/01-edu/z01"
)

func main() {
	table := append(
		z01.MultRandASCII(),
		"Hello! How are you? How+are+things+4you?",
		"Hello! How are you?",
		"a",
		"z",
		"!",
		"9a",
		"9a LALALA!",
	)
	for _, arg := range table {
		z01.ChallengeMain(arg)
	}
	z01.ChallengeMain("hello", "hihihi")
	z01.ChallengeMain()
}
