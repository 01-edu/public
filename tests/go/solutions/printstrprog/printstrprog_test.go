package main

import (
	"testing"

	"github.com/01-edu/z01"
)

func TestPrintStrProg(t *testing.T) {
	table := z01.MultRandASCII()
	for _, arg := range table {
		z01.ChallengeMainExam(t, arg)
	}
	z01.ChallengeMainExam(t, "Hello World!")
}
