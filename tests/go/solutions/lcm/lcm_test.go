package main

import (
	"github.com/01-edu/z01"

	solutions "../../solutions"
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
			z01.RandIntBetween(0, 1000),
			z01.RandIntBetween(0, 1000),
		})
	}

	for _, arg := range table {
		z01.Challenge("Lcm", Lcm, solutions.Lcm, arg[0], arg[1])
	}
}
