package student_test

import (
	"testing"

	solutions "./solutions"
	student "./student"
	"github.com/01-edu/z01"
)

func TestCollatzCountdown(t *testing.T) {
	args := []int{z01.RandIntBetween(-6, 20)}
	args = append(args, -5, 0)
	for i := 0; i < 20; i++ {
		args = append(args, z01.RandIntBetween(2, 20))
	}

	for _, v := range args {
		z01.Challenge(t, student.CollatzCountdown, solutions.CollatzCountdown, v)
	}
}
