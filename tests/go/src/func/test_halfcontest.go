package main

import (
	"../lib"
	"./correct"
	"./student"
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
			h1: lib.RandIntBetween(0, 10),
			m1: lib.RandIntBetween(0, 59),
			h2: lib.RandIntBetween(11, 23),
			m2: lib.RandIntBetween(0, 59),
		})
	}

	for _, arg := range table {
		lib.Challenge("HalfContest", student.HalfContest, correct.HalfContest, arg.h1, arg.m1, arg.h2, arg.m2)
	}
}
