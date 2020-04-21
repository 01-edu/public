package main

import (
	"github.com/01-edu/z01"
)

func main() {
	table := append(z01.MultRandWords(), " ")
	table = append(table, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPRSTUVWXYZ")
	table = append(table, "a b c d e f g h ijklmnopqrstuvwxyz A B C D E FGHIJKLMNOPRSTUVWXYZ")

	for _, s := range table {
		z01.ChallengeMain("rot13", s)
	}
	z01.ChallengeMain("rot13", "1 argument", "2 arguments")
	z01.ChallengeMain("rot13", "1 argument", "2 arguments", "3 arguments")
}
