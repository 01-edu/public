package student_test

import (
	"testing"

	"github.com/01-edu/z01"

	solutions "./solutions"
	student "./student"
)

func TestPrintStr(t *testing.T) {
	table := z01.MultRandASCII()
	table = append(table, "Hello World!")
	for _, arg := range table {
		z01.Challenge(t, student.PrintStr, solutions.PrintStr, arg)
	}
}
