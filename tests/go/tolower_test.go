package student_test

import (
	"testing"

	"github.com/01-edu/z01"

	solutions "./solutions"
	student "./student"
)

func TestToLower(t *testing.T) {
	table := append(
		z01.MultRandASCII(),
		"Hello! How are you?",
	)
	for _, arg := range table {
		z01.Challenge(t, student.ToLower, solutions.ToLower, arg)
	}
}
