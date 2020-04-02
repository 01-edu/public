package student_test

import (
	"testing"

	"github.com/01-edu/z01"

	solutions "./solutions"
	student "./student"
)

func TestItoaBase(t *testing.T) {
	for i := 0; i < 30; i++ {
		value := z01.RandIntBetween(-1000000, 1000000)
		base := z01.RandIntBetween(2, 16)
		z01.Challenge(t, student.ItoaBase, solutions.ItoaBase, value, base)
	}
	for i := 0; i < 5; i++ {
		base := z01.RandIntBetween(2, 16)
		z01.Challenge(t, student.ItoaBase, solutions.ItoaBase, z01.MaxInt, base)
		z01.Challenge(t, student.ItoaBase, solutions.ItoaBase, z01.MinInt, base)
	}
}
