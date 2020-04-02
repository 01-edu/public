package student_test

import (
	"testing"

	"github.com/01-edu/z01"

	solutions "./solutions"
	student "./student"
)

func TestFibonacci(t *testing.T) {
	table := append(
		z01.MultRandIntBetween(0, 25),
		4,
		5,
		-5,
	)
	for _, arg := range table {
		z01.Challenge(t, student.Fibonacci, solutions.Fibonacci, arg)
	}
}
