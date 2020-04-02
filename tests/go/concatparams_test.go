package student_test

import (
	"testing"

	"github.com/01-edu/z01"

	solutions "./solutions"
	student "./student"
)

func TestConcatParams(t *testing.T) {

	table := [][]string{}
	//30 random slice of strings

	for i := 0; i < 30; i++ {
		val := z01.MultRandASCII()
		table = append(table, val)
	}

	table = append(table,
		[]string{"Hello", "how", "are", "you?"},
	)

	for _, arg := range table {
		z01.Challenge(t, student.ConcatParams, solutions.ConcatParams, arg)
	}
}
