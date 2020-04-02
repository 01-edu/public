package student_test

import (
	"testing"

	"github.com/01-edu/z01"

	solutions "./solutions"
	student "./student"
)

func TestIsAlpha(t *testing.T) {
	table := append(
		z01.MultRandASCII(),
		"Hello! €How are you?",
		"a",
		"z",
		"!",
		"HelloHowareyou",
		"What's this 4?",
		"Whatsthis4",
		"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890",
		"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!",
	)
	for _, arg := range table {
		z01.Challenge(t, student.IsAlpha, solutions.IsAlpha, arg)
	}
}
