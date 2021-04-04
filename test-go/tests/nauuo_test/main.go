package main

import (
	student "student"

	"github.com/01-edu/public/test-go/lib"
)

func nauuo(plus, minus, rand int) string {
	if rand == 0 {
		if plus > minus {
			return "+"
		}
		if plus < minus {
			return "-"
		}
		if plus == minus {
			return "0"
		}
	}
	if plus > minus+rand {
		return "+"
	}
	if plus+rand < minus {
		return "-"
	}
	if plus+rand >= minus && plus-rand <= minus {
		return "?"
	}
	if minus+rand >= plus && minus-rand <= plus {
		return "?"
	}
	return "?"
}

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
		lib.Challenge("Nauuo", student.Nauuo, nauuo, arg.plus, arg.minus, arg.rand)
	}
}
