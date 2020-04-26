package main

import (
	"github.com/01-edu/z01"

	correct "./correct"
	student "./student"
)

func main() {
	type node struct {
		red   int
		green int
		blue  int
	}

	table := []node{
		{50, 43, 20},
		{10, 0, 0},
		{0, 10, 0},
		{0, 0, 10},
		{10, 1, 0},
		{0, 10, 1},
		{1, 0, 10},
		{10, 2, 0},
		{2, 10, 0},
		{0, 2, 10},
		{13, 13, 0},
		{10, 9, 0},
		{5, 9, 2},
	}

	for i := 0; i < 15; i++ {
		table = append(table, node{
			red:   z01.RandIntBetween(0, 30),
			green: z01.RandIntBetween(0, 30),
			blue:  z01.RandIntBetween(0, 30),
		})
	}
	for _, arg := range table {
		z01.Challenge("SweetProblem", student.Sweetproblem, correct.Sweetproblem, arg.red, arg.green, arg.blue)
	}
}
