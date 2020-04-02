package main

import (
	"strings"
	"testing"

	"github.com/01-edu/z01"
)

func TestDisplayz(t *testing.T) {
	table := append(z01.MultRandWords(), "dsfdz")
	table = append(table, "")
	table = append(table, "1")
	table = append(table, "1")

	for _, s := range table {
		z01.ChallengeMainExam(t, strings.Fields(s)...)
	}

	z01.ChallengeMainExam(t, "1", "z")
}
