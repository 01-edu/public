package main

import (
	"github.com/01-edu/z01"

	correct "./correct"
	student "./student"
)

func main() {
	type node struct {
		h1, m1 int
		h2, m2 int
	}
	table := []node{
		{11, 44, 21, 59},
		{1, 12, 1, 14},
		{5, 50, 6, 51},
		{14, 35, 18, 55},
	}

	for i := 0; i < 20; i++ {
		table = append(table, node{
			h1: z01.RandIntBetween(0, 10),
			m1: z01.RandIntBetween(0, 59),
			h2: z01.RandIntBetween(11, 23),
			m2: z01.RandIntBetween(0, 59),
		})
	}

	for _, arg := range table {
		z01.Challenge("HalfContestProg", student.Halfcontest, correct.Halfcontest, arg.h1, arg.m1, arg.h2, arg.m2)
	}
}
