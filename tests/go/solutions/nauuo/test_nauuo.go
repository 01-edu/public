package main

import (
	"github.com/01-edu/z01"

	correct "./correct"
)

func main() {
	type node struct {
		plus  int
		minus int
		rand  int
	}
	table := []node{
		{50, 43, 20},
		{13, 13, 0},
		{10, 9, 0},
		{5, 9, 2},
	}

	for i := 0; i < 15; i++ {
		table = append(table, node{
			plus:  z01.RandIntBetween(0, 10),
			minus: z01.RandIntBetween(0, 10),
			rand:  z01.RandIntBetween(0, 10),
		})
	}
	for _, arg := range table {
		z01.Challenge("Nauuo", Nauuo, correct.Nauuo, arg.plus, arg.minus, arg.rand)
	}
}
