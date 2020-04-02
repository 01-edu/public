package main

import (
	"strings"
	"testing"

	"github.com/01-edu/z01"
)

func TestWdMatch(t *testing.T) {
	table := append(z01.MultRandWords(),
		" ",
		"faya fgvvfdxcacpolhyghbreda",
		"faya fgvvfdxcacpolhyghbred",
		"error rrerrrfiiljdfxjyuifrrvcoojh",
		"error rrerrrfiiljdfxjyuifrrvcoojrh",
	)

	for _, s := range table {
		z01.ChallengeMainExam(t, strings.Fields(s)...)
	}
}
