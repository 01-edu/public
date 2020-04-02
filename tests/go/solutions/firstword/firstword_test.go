package main

import (
	"testing"

	"github.com/01-edu/z01"
)

func TestFirstWord(t *testing.T) {
	table := append(z01.MultRandWords(),
		"",
		"             a as",
		"   f     d",
		"   asd    ad",
	)
	table = append(table, "   salut !!! ")
	table = append(table, " salut ! ! !")
	table = append(table, "salut ! !")

	for _, s := range table {
		z01.ChallengeMainExam(t, s)
	}
}
