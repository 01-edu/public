package main

import (
	"../lib"
	"./correct"
	"./student"
)

func main() {
	functions := []func(string) bool{correct.IsNumeric, correct.IsLower, correct.IsUpper}

	type node struct {
		f   func(string) bool
		arr []string
	}

	table := []node{}

	for i := 0; i < 5; i++ {
		function := functions[lib.RandIntBetween(0, len(functions)-1)]
		val := node{
			f:   function,
			arr: lib.MultRandWords(),
		}
		table = append(table, val)
	}
	for i := 0; i < 5; i++ {
		val := node{
			f:   correct.IsNumeric,
			arr: lib.MultRandDigit(),
		}
		table = append(table, val)
	}

	for i := 0; i < 5; i++ {
		val := node{
			f:   correct.IsLower,
			arr: lib.MultRandLower(),
		}
		table = append(table, val)
	}
	for i := 0; i < 5; i++ {
		val := node{
			f:   correct.IsUpper,
			arr: lib.MultRandUpper(),
		}
		table = append(table, val)
	}

	table = append(table,
		node{
			f:   correct.IsNumeric,
			arr: []string{"Hello", "how", "are", "you"},
		},
		node{
			f:   correct.IsNumeric,
			arr: []string{"This", "is", "4", "you"},
		},
	)

	for _, arg := range table {
		lib.Challenge("CountIf", student.CountIf, correct.CountIf, arg.f, arg.arr)
	}
}
