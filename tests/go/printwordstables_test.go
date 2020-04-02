package student_test

import (
	"strings"
	"testing"

	"github.com/01-edu/z01"

	solutions "./solutions"
	student "./student"
)

func TestPrintWordsTables(t *testing.T) {

	table := [][]string{}
	//30 random slice of slice of strings

	for i := 0; i < 30; i++ {

		val := solutions.SplitWhiteSpaces(strings.Join(z01.MultRandASCII(), " "))
		table = append(table, val)
	}

	table = append(table,
		[]string{"Hello", "how", "are", "you?"})

	for _, arg := range table {
		z01.Challenge(t, student.PrintWordsTables, solutions.PrintWordsTables, arg)
	}
}
