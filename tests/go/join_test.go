package student_test

import (
	solutions "./solutions"
	student "./student"
	"github.com/01-edu/z01"
	"testing"
)

func TestJoin(t *testing.T) {
	arg1 := []string{"hello", "how", "are", "you", "doing"}
	arg2 := []string{"fine", "and", "you"}
	arg3 := []string{"I'm", "O.K."}
	seps := []string{" ", "-", " ,", "_", "SPC", " . "}

	args := [][]string{arg1, arg2, arg3}

	for i := 0; i < 5; i++ {
		//random position for the array of arguments
		posA := z01.RandIntBetween(0, len(args)-1)
		//random position for the array of separators
		posS := z01.RandIntBetween(0, len(seps)-1)

		z01.Challenge(t, student.Join, solutions.Join, args[posA], seps[posS])
	}
}
