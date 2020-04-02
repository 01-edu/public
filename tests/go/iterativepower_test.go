package student_test

import (
	"testing"

	"github.com/01-edu/z01"

	solutions "./solutions"
	student "./student"
)

func TestIterativePower(t *testing.T) {
	i := 0
	for i < 30 {
		nb := z01.RandIntBetween(-8, 8)
		power := z01.RandIntBetween(-10, 10)
		z01.Challenge(t, student.IterativePower, solutions.IterativePower, nb, power)
		i++
	}
	z01.Challenge(t, student.IterativePower, solutions.IterativePower, 0, 0)
	z01.Challenge(t, student.IterativePower, solutions.IterativePower, 0, 1)
}
