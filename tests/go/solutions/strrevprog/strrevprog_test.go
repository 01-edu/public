package main

import (
	"testing"

	"github.com/01-edu/z01"
)

func TestStrRevProg(t *testing.T) {
	table := append(
		z01.MultRandASCII(),
		"Hello!",
		"Bonjour!",
		"Hola!",
	)
	for _, arg := range table {
		z01.ChallengeMainExam(t, arg)
	}
}
