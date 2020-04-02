package student_test

import (
	"testing"

	"github.com/01-edu/z01"

	solutions "./solutions"
	student "./student"
)

func TestCapitalize(t *testing.T) {
	table := append(
		z01.MultRandASCII(),
		"Hello! How are you? How+are+things+4you?",
		"Hello! How are you?",
		"a",
		"z",
		"!",
		"9a",
		"9a LALALA!",
	)
	for _, arg := range table {
		z01.Challenge(t, student.Capitalize, solutions.Capitalize, arg)
	}
}
