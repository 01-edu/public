package main

import (
	"../lib"
	"./correct"
	"./student"
)

func main() {
	table := [][2]int{
		{50, 43},
		{13, 13},
		{10, 9},
		{5, 9},
	}
	for i := 0; i < 15; i++ {
		table = append(table, [2]int{
			lib.RandIntBetween(0, 30),
			lib.RandIntBetween(0, 100),
		})
	}
	for _, arg := range table {
		lib.Challenge("Volumechanger", student.Volumechanger, correct.Volumechanger, arg[0], arg[1])
	}
}
