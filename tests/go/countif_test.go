package student_test

import (
	"testing"

	"github.com/01-edu/z01"

	solutions "./solutions"
	student "./student"
)

func TestCountIf(t *testing.T) {
	functionsArray := []func(string) bool{solutions.IsNumeric, solutions.IsLower, solutions.IsUpper}

	type node struct {
		f   func(string) bool
		arr []string
	}

	table := []node{}

	for i := 0; i < 5; i++ {

		functionSelected := functionsArray[z01.RandIntBetween(0, len(functionsArray)-1)]
		val := node{
			f:   functionSelected,
			arr: z01.MultRandWords(),
		}
		table = append(table, val)

	}
	for i := 0; i < 5; i++ {

		val := node{
			f:   solutions.IsNumeric,
			arr: z01.MultRandDigit(),
		}
		table = append(table, val)

	}

	for i := 0; i < 5; i++ {

		val := node{
			f:   solutions.IsLower,
			arr: z01.MultRandLower(),
		}
		table = append(table, val)

	}
	for i := 0; i < 5; i++ {

		val := node{
			f:   solutions.IsUpper,
			arr: z01.MultRandUpper(),
		}
		table = append(table, val)

	}

	table = append(table,
		node{
			f:   solutions.IsNumeric,
			arr: []string{"Hello", "how", "are", "you"},
		},
		node{
			f:   solutions.IsNumeric,
			arr: []string{"This", "is", "4", "you"},
		},
	)

	for _, arg := range table {
		z01.Challenge(t, student.CountIf, solutions.CountIf, arg.f, arg.arr)
	}
}
