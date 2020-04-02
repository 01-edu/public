package student_test

import (
	"testing"

	"github.com/01-edu/z01"

	solutions "./solutions"
	student "./student"
)

func TestPrintNbrBase(t *testing.T) {
	type node struct {
		n    int
		base string
	}

	table := []node{}

	// 15 random pairs of ints with valid bases
	for i := 0; i < 15; i++ {
		validBaseToInput := solutions.RandomValidBase()
		val := node{
			n:    z01.RandIntBetween(-1000000, 1000000),
			base: validBaseToInput,
		}
		table = append(table, val)
	}

	// 15 random pairs of ints with invalid bases
	for i := 0; i < 15; i++ {
		invalidBaseToInput := solutions.RandomInvalidBase()
		val := node{
			n:    z01.RandIntBetween(-1000000, 1000000),
			base: invalidBaseToInput,
		}
		table = append(table, val)
	}

	table = append(table,
		node{n: 125, base: "0123456789"},
		node{n: -125, base: "01"},
		node{n: 125, base: "0123456789ABCDEF"},
		node{n: -125, base: "choumi"},
		node{n: 125, base: "-ab"},
		node{n: z01.MinInt, base: "0123456789"},
	)
	for _, arg := range table {
		z01.Challenge(t, student.PrintNbrBase, solutions.PrintNbrBase, arg.n, arg.base)
	}
}
