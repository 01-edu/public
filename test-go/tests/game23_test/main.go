package main

import (
	student "student"

	"github.com/01-edu/public/test-go/lib"
)

func game23(a, b int) int {
	if a > b {
		return -1
	}
	if a == b {
		return 0
	}
	if game23(a*2, b) != -1 {
		return 1 + game23(a*2, b)
	}
	if game23(a*3, b) != -1 {
		return 1 + game23(a*3, b)
	}
	return -1
}

func nd(a, b int) int {
	if a > b {
		return -1
	}
	if a == b {
		return 0
	}
	if nd(a*2, b) != -1 {
		return 1 + nd(a*2, b)
	}
	if nd(a*3, b) != -1 {
		return 1 + nd(a*3, b)
	}
	return -1
}

func main() {
	type node struct {
		init int
		fin  int
	}
	table := []node{
		{50, 43},
		{13, 13},
		{10, 9},
		{5, 9},
	}

	for i := 0; i < 20; i++ {
		table = append(table, node{
			init: lib.RandIntBetween(1, 1000),
			fin:  lib.RandIntBetween(1, 1000),
		})
	}

	for i := 1; i < 100; i++ {
		value := node{
			init: 1,
			fin:  1,
		}
		if nd(1, i) > 0 {
			value = node{
				init: 1,
				fin:  i,
			}
		}
		table = append(table, value)
	}
	for _, arg := range table {
		lib.Challenge("Game23", student.Game23, game23, arg.init, arg.fin)
	}
}
