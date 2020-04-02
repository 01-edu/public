package student_test

import (
	"testing"

	"github.com/01-edu/z01"

	solutions "./solutions"
	student "./student"
)

func TestToUpper(t *testing.T) {

	table := z01.MultRandASCII()

	table = append(table,
		"Hello! How are you?",
	)

	for _, arg := range table {
		z01.Challenge(t, student.ToUpper, solutions.ToUpper, arg)
	}
}
