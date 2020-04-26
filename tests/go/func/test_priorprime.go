package main

import (
	"../lib"
	"./correct"
	"./student"
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
	table = append(table, lib.MultRandIntBetween(0, 1000))
	for _, arg := range table {
		lib.Challenge("PriorPrime", student.PriorPrime, correct.PriorPrime, arg)
	}
}
