package main

import (
	"github.com/01-edu/z01"

	correct "./correct"
)

func main() {
	table := []int{
		20,
		0,
		9,
		2,
	}
	table = append(table, z01.MultRandIntBetween(-100, 150)...)
	for _, arg := range table {
		z01.Challenge("Fib", Fib, correct.Fib, arg)
	}
}
