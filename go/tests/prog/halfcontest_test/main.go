package main

import (
	student "student"

	"github.com/01-edu/public/go/tests/lib"
)

func halfContest(h1, m1, h2, m2 int) int {
	t1 := h1*60 + m1
	t2 := h2*60 + m2
	t2 = (t2 + t1) / 2
	h2 = t2 / 60
	m2 = t2 % 60
	return h2*100 + m2
}

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
		lib.Challenge("HalfContest", student.HalfContest, halfContest, arg.h1, arg.m1, arg.h2, arg.m2)
	}
}
