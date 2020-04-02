package student_test

import (
	"testing"

	"github.com/01-edu/z01"

	solutions "./solutions"
	student "./student"
)

func TestIsNegative(t *testing.T) {
	table := append(
		z01.MultRandInt(),
		z01.MinInt,
		z01.MaxInt,
		0,
	)
	for _, arg := range table {
		z01.Challenge(t, student.IsNegative, solutions.IsNegative, arg)
	}
}
