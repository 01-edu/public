package student_test

import (
	"testing"

	"github.com/01-edu/z01"

	solutions "./solutions"
	student "./student"
)

func TestPrintCombN(t *testing.T) {
	table := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, arg := range table {
		z01.Challenge(t, student.PrintCombN, solutions.PrintCombN, arg)
	}
}
