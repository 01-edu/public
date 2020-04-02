package student_test

import (
	"testing"

	"github.com/01-edu/z01"

	solutions "./solutions"
	student "./student"
)

func TestConvertBase(t *testing.T) {

	type node struct {
		nbr      string
		baseFrom string
		baseTo   string
	}
	table := []node{}

	for i := 0; i < 30; i++ {
		validBaseToInput1 := solutions.RandomValidBase()
		validBaseToInput2 := solutions.RandomValidBase()
		str := solutions.ConvertNbrBase(z01.RandIntBetween(0, 1000000), validBaseToInput1)
		val := node{
			nbr:      str,
			baseFrom: validBaseToInput1,
			baseTo:   validBaseToInput2,
		}
		table = append(table, val)
	}

	table = append(table, node{
		nbr:      "101011",
		baseFrom: "01",
		baseTo:   "0123456789"},
	)

	for _, arg := range table {
		z01.Challenge(t, student.ConvertBase, solutions.ConvertBase, arg.nbr, arg.baseFrom, arg.baseTo)
	}
}
