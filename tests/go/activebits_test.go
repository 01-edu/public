package student_test

import (
	solutions "./solutions"
	student "./student"
	"github.com/01-edu/z01"
	"testing"
)

func TestActiveBits(t *testing.T) {
	args := []int{z01.RandIntBetween(2, 20)}

	for i := 0; i < 20; i++ {
		args = append(args, z01.RandIntBetween(2, 20))
	}

	for _, v := range args {
		z01.Challenge(t, student.ActiveBits, solutions.ActiveBits, v)
	}

}
