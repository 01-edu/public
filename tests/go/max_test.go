package student_test

import (
	solutions "./solutions"
	student "./student"
	"github.com/01-edu/z01"
	"testing"
)

func TestMax(t *testing.T) {
	args := []int{z01.RandInt()}
	limit := z01.RandIntBetween(20, 50)
	for i := 0; i < limit; i++ {
		args = append(args, z01.RandInt())
	}

	z01.Challenge(t, student.Max, solutions.Max, args)
}
