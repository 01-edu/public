package main

import (
	"../lib"
	"./correct"
	"./student"
)

func main() {
	table := []int{
		20,
		0,
		9,
		2,
	}
	table = append(table, lib.MultRandIntBetween(-100, 150)...)
	for _, arg := range table {
		lib.Challenge("Fib", student.Fib, correct.Fib, arg)
	}
}
