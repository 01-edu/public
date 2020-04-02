package student_test

import (
	"testing"

	"github.com/01-edu/z01"

	solutions "./solutions"
	student "./student"
)

func TestSqrt(t *testing.T) {
	table := append(
		z01.MultRandIntBetween(-1000000, 1000000),
		0,
		1,
		2,
		3,
		4,
		5,
		6,
		7,
		8,
		9,
		10,
		11,
		12,
		100,
	)
	for _, arg := range table {
		z01.Challenge(t, student.Sqrt, solutions.Sqrt, arg)
	}
}
