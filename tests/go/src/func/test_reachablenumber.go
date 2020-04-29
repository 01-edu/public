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
	for i := 0; i < 25; i++ {
		table = append(table, lib.MultRandIntBetween(1, 877))
	}
	for _, arg := range table {
		lib.Challenge("ReachableNumber", student.ReachableNumber, correct.ReachableNumber, arg)
	}
}
