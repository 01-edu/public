package main

import (
	"../common"
	"./student"
)

func countIf(f func(string) bool, arr []string) int {
	counter := 0
	for _, el := range arr {
		if f(el) {
			counter++
		}
	}

	return counter
}

func main() {
	functions := []func(string) bool{common.IsNumeric, common.IsLower, common.IsUpper}

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
			f:   common.IsNumeric,
			arr: lib.MultRandDigit(),
		}
		table = append(table, val)
	}

	for i := 0; i < 5; i++ {
		val := node{
			f:   common.IsLower,
			arr: lib.MultRandLower(),
		}
		table = append(table, val)
	}
	for i := 0; i < 5; i++ {
		val := node{
			f:   common.IsUpper,
			arr: lib.MultRandUpper(),
		}
		table = append(table, val)
	}

	table = append(table,
		node{
			f:   common.IsNumeric,
			arr: []string{"Hello", "how", "are", "you"},
		},
		node{
			f:   common.IsNumeric,
			arr: []string{"This", "is", "4", "you"},
		},
	)

	for _, arg := range table {
		lib.Challenge("CountIf", student.CountIf, countIf, arg.f, arg.arr)
	}
}
