package main

import (
	"../lib"
	"./correct"
	"./student"
)

func main() {
	functions := []func(int) bool{correct.IsPositive, correct.IsNegative0, correct.IsPrime}

	type node struct {
		f   func(int) bool
		arr []int
	}

	table := []node{}

	for i := 0; i < 15; i++ {
		function := functions[lib.RandIntBetween(0, len(functions)-1)]
		val := node{
			f:   function,
			arr: lib.MultRandIntBetween(-1000000, 1000000),
		}
		table = append(table, val)
	}

	table = append(table, node{
		f:   correct.IsPrime,
		arr: []int{1, 2, 3, 4, 5, 6},
	})

	for _, arg := range table {
		lib.Challenge("Map", student.Map, correct.Map, arg.f, arg.arr)
	}
}
