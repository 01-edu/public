package main

import (
	"github.com/01-edu/z01"

	correct "./correct"
)

func main() {
	type node struct {
		red   int
		green int
		blue  int
	}

	table := []node{}

	table = append(table,
		node{50, 43, 20},
		node{10, 0, 0},
		node{0, 10, 0},
		node{0, 0, 10},
		node{10, 1, 0},
		node{0, 10, 1},
		node{1, 0, 10},
		node{10, 2, 0},
		node{2, 10, 0},
		node{0, 2, 10},
		node{13, 13, 0},
		node{10, 9, 0},
		node{5, 9, 2},
	)

	for i := 0; i < 15; i++ {
		value := node{
			red:   z01.RandIntBetween(0, 30),
			green: z01.RandIntBetween(0, 30),
			blue:  z01.RandIntBetween(0, 30),
		}
		table = append(table, value)
	}
	for _, arg := range table {
		z01.Challenge("SweetProblem", Sweetproblem, correct.Sweetproblem, arg.red, arg.green, arg.blue)
	}
}
