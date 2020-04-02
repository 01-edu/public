package student_test

import (
	"testing"

	"github.com/01-edu/z01"

	solutions "./solutions"
	student "./student"
)

func TestMakeRange(t *testing.T) {
	type node struct {
		min int
		max int
	}

	table := []node{}

	// 15 random pairs of ints for a Valid Range
	for i := 0; i < 15; i++ {
		minVal := z01.RandIntBetween(-10000000, 1000000)
		gap := z01.RandIntBetween(1, 20)
		val := node{
			min: minVal,
			max: minVal + gap,
		}
		table = append(table, val)
	}
	// 15 random pairs of ints with ||invalid range||
	for i := 0; i < 15; i++ {
		minVal := z01.RandIntBetween(-10000000, 1000000)
		gap := z01.RandIntBetween(1, 20)
		val := node{
			min: minVal,
			max: minVal - gap,
		}
		table = append(table, val)
	}

	table = append(table,
		node{min: 0, max: 1},
		node{min: 0, max: 0},
		node{min: 5, max: 10},
		node{min: 10, max: 5},
	)
	for _, arg := range table {
		z01.Challenge(t, student.MakeRange, solutions.MakeRange, arg.min, arg.max)
	}
}
