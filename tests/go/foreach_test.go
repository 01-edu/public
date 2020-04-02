package student_test

import (
	"testing"

	"github.com/01-edu/z01"

	solutions "./solutions"
	student "./student"
)

func TestForEach(t *testing.T) {

	functionsArray := []func(int){solutions.Add0, solutions.Add1, solutions.Add2, solutions.Add3}

	type node struct {
		f   func(int)
		arr []int
	}

	table := []node{}

	//15 random slice of random ints with a random function from selection

	for i := 0; i < 15; i++ {

		functionSelected := functionsArray[z01.RandIntBetween(0, len(functionsArray)-1)]
		val := node{
			f:   functionSelected,
			arr: z01.MultRandIntBetween(-1000000, 1000000),
		}
		table = append(table, val)

	}

	table = append(table, node{
		f:   solutions.Add0,
		arr: []int{1, 2, 3, 4, 5, 6},
	})

	for _, arg := range table {
		z01.Challenge(t, student.ForEach, solutions.ForEach, arg.f, arg.arr)
	}
}
