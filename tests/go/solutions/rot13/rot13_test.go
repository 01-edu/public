package main

import (
	"testing"

	"github.com/01-edu/z01"
)

func TestRot13(t *testing.T) {
	table := append(z01.MultRandWords(), " ")
	table = append(table, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPRSTUVWXYZ")
	table = append(table, "a b c d e f g h ijklmnopqrstuvwxyz A B C D E FGHIJKLMNOPRSTUVWXYZ")

	for _, s := range table {
		z01.ChallengeMainExam(t, s)
	}
	z01.ChallengeMainExam(t, "1 argument", "2 arguments")
	z01.ChallengeMainExam(t, "1 argument", "2 arguments", "3 arguments")
}
