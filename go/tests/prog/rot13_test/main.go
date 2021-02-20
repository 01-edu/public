package main

import (
	"github.com/01-edu/public/go/tests/lib"
)

func main() {
	table := append(lib.MultRandWords(),
		" ",
		"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPRSTUVWXYZ",
		"a b c d e f g h ijklmnopqrstuvwxyz A B C D E FGHIJKLMNOPRSTUVWXYZ",
	)
	for _, s := range table {
		lib.ChallengeMain("rot13", s)
	}
	lib.ChallengeMain("rot13", "1 argument", "2 arguments")
	lib.ChallengeMain("rot13", "1 argument", "2 arguments", "3 arguments")
}
