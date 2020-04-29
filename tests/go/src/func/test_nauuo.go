package main

import (
	"../lib"
	"./correct"
	"./student"
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
			plus:  lib.RandIntBetween(0, 10),
			minus: lib.RandIntBetween(0, 10),
			rand:  lib.RandIntBetween(0, 10),
		})
	}
	for _, arg := range table {
		lib.Challenge("Nauuo", student.Nauuo, correct.Nauuo, arg.plus, arg.minus, arg.rand)
	}
}
