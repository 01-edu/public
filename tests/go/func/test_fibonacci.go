package main

import (
	"../lib"
	"./correct"
	"./student"
)

func main() {
	table := append(
		lib.MultRandIntBetween(0, 25),
		4,
		5,
		-5,
	)
	for _, arg := range table {
		lib.Challenge("Fibonacci", student.Fibonacci, correct.Fibonacci, arg)
	}
}
