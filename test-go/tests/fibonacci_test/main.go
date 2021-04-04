package main

import (
	student "student"

	"github.com/01-edu/public/test-go/lib"
)

func fibonacci(i int) int {
	if i < 0 {
		return -1
	}
	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}
	return fibonacci(i-1) + fibonacci(i-2)
}

func main() {
	table := append(lib.MultRandIntBetween(0, 25),
		4,
		5,
		-5,
	)
	for _, arg := range table {
		lib.Challenge("Fibonacci", student.Fibonacci, fibonacci, arg)
	}
}
