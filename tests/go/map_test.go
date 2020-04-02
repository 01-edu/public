package student_test

import (
	"testing"

	"github.com/01-edu/z01"

	solutions "./solutions"
	student "./student"
)

func TestMap(t *testing.T) {
	functionsArray := []func(int) bool{solutions.IsPositive, solutions.IsNegative0, solutions.IsPrime}

	type node struct {
		f   func(int) bool
		arr []int
	}

	table := []node{}

	for i := 0; i < 15; i++ {

		functionSelected := functionsArray[z01.RandIntBetween(0, len(functionsArray)-1)]
		val := node{
			f:   functionSelected,
			arr: z01.MultRandIntBetween(-1000000, 1000000),
		}
		table = append(table, val)

	}

	table = append(table, node{
		f:   solutions.IsPrime,
		arr: []int{1, 2, 3, 4, 5, 6},
	})

	for _, arg := range table {
		z01.Challenge(t, student.Map, solutions.Map, arg.f, arg.arr)
	}
}
