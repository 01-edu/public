package student_test

import (
	"testing"

	solutions "./solutions"
	student "./student"
	"github.com/01-edu/z01"
)

func TestAbort(t *testing.T) {
	arg := z01.MultRandInt()
	arg = append(arg, z01.RandInt())
	for i := 0; i < 15; i++ {
		z01.Challenge(t, student.Abort, solutions.Abort, arg[0], arg[1], arg[2], arg[3], arg[4])
		arg = z01.MultRandInt()
		arg = append(arg, z01.RandInt())
	}
}
