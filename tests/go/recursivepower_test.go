package student_test

import (
	"testing"

	"github.com/01-edu/z01"

	solutions "./solutions"
	student "./student"
)

func TestRecursivePower(t *testing.T) {
	i := 0
	for i < 30 {
		nb := z01.RandIntBetween(-8, 8)
		power := z01.RandIntBetween(-10, 10)
		z01.Challenge(t, student.RecursivePower, solutions.RecursivePower, nb, power)
		i++
	}
	z01.Challenge(t, student.RecursivePower, solutions.RecursivePower, 0, 0)
	z01.Challenge(t, student.RecursivePower, solutions.RecursivePower, 0, 1)
}
