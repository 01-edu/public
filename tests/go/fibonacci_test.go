package main

import (
	"github.com/01-edu/z01"

	solutions "./solutions"
	student "./student"
)

func main() {
	table := append(
		z01.MultRandIntBetween(0, 25),
		4,
		5,
		-5,
	)
	for _, arg := range table {
		z01.Challenge("Fibonacci", student.Fibonacci, solutions.Fibonacci, arg)
	}
}
