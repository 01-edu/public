package main

import (
	"../lib"
	"./correct"
	"./student"
)

func main() {
	table := []int{
		20,
		1,
		9,
		2,
	}
	table = append(table, lib.MultRandIntBetween(1, 1500)...)

	for _, arg := range table {
		lib.Challenge("InterestingNumber", student.InterestingNumber, correct.InterestingNumber, arg)
	}
}
