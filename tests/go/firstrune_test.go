package student_test

import (
	"testing"

	"github.com/01-edu/z01"

	solutions "./solutions"
	student "./student"
)

func TestFirstRune(t *testing.T) {
	table := append(
		z01.MultRandASCII(),
		"Hello!",
		"Salut!",
		"Ola!",
		"♥01",
	)
	for _, arg := range table {
		z01.Challenge(t, student.FirstRune, solutions.FirstRune, arg)
	}
}
