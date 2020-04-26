package main

import (
	"github.com/01-edu/z01"

	solutions "../../solutions"
)

func main() {
	table := []int{
		50,
		13,
		10,
		0,
		1,
		2,
	}
	table = append(table, z01.MultRandIntBetween(0, 1000))
	for _, arg := range table {
		z01.Challenge("PriorPrime", PriorPrime, solutions.PriorPrime, arg)
	}
}
