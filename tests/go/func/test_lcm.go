package main

import (
	"../lib"
	"./correct"
	"./student"
)

func main() {
	table := [][2]int{
		{50, 43},
		{13, 13},
		{10, 9},
		{0, 9},
		{1, 1},
	}

	for i := 0; i < 15; i++ {
		table = append(table, [2]int{
			lib.RandIntBetween(0, 1000),
			lib.RandIntBetween(0, 1000),
		})
	}

	for _, arg := range table {
		lib.Challenge("Lcm", student.Lcm, correct.Lcm, arg[0], arg[1])
	}
}
