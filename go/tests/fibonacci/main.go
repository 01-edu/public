package main

import (
	"./student"
	"github.com/01-edu/public/go/lib"
)

func fibonacci(value int) int {
	if value < 0 {
		return -1
	}
	if value == 0 {
		return 0
	}
	if value == 1 {
		return 1
	}
	return Fibonacci(value-1) + Fibonacci(value-2)
}

func main() {
	table := append(
		lib.MultRandIntBetween(0, 25),
		4,
		5,
		-5,
	)
	for _, arg := range table {
		lib.Challenge("Fibonacci", student.Fibonacci, fibonacci, arg)
	}
}
