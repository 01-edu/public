package main

import (
	"strings"
	"testing"

	"github.com/01-edu/z01"
)

func TestDisplayLastParam(t *testing.T) {
	table := append(z01.MultRandWords(), " ")
	table = append(table, "1")
	table = append(table, "1 2")

	for _, s := range table {
		z01.ChallengeMainExam(t, strings.Fields(s)...)
	}
}
