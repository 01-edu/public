package student_test

import (
	"testing"

	"github.com/01-edu/z01"

	solutions "./solutions"
	student "./student"
)

func TestStrRev(t *testing.T) {
	table := append(
		z01.MultRandASCII(),
		"Hello!",
		"Bonjour!",
		"Hola!",
	)
	for _, arg := range table {
		z01.Challenge(t, student.StrRev, solutions.StrRev, arg)
	}
}
