package main

import (
	"../lib"
	"./correct"
	"./student"
)

func main() {
	var table [10]int

	for j := 0; j < 5; j++ {
		for i := 0; i < 10; i++ {
			table[i] = lib.RandIntBetween(0, 1000)
		}
		lib.Challenge("PrintMemory", student.PrintMemory, correct.PrintMemory, table)
	}
	table2 := [10]int{104, 101, 108, 108, 111, 16, 21, 42}
	lib.Challenge("PrintMemory", student.PrintMemory, correct.PrintMemory, table2)
}

// TODO: this can be simplified a lot
