package main

import (
	"../lib"
	"./correct"
	"./student"
)

func main() {
	functions := []func(string) bool{correct.IsNumeric, correct.IsLower, correct.IsUpper}

	type node struct {
		f func(string) bool
		a []string
	}

	table := []node{}

	for i := 0; i < 5; i++ {
		function := functions[lib.RandIntBetween(0, len(functions)-1)]
		table = append(table, node{
			f: function,
			a: lib.MultRandWords(),
		})
	}
	for i := 0; i < 5; i++ {
		table = append(table, node{
			f: correct.IsNumeric,
			a: lib.MultRandDigit(),
		})
	}

	for i := 0; i < 5; i++ {
		table = append(table, node{
			f: correct.IsLower,
			a: lib.MultRandLower(),
		})
	}
	for i := 0; i < 5; i++ {
		table = append(table, node{
			f: correct.IsUpper,
			a: lib.MultRandUpper(),
		})
	}

	table = append(table,
		node{
			f: correct.IsNumeric,
			a: []string{"Hello", "how", "are", "you"},
		},
		node{
			f: correct.IsNumeric,
			a: []string{"This", "is", "4", "you"},
		},
	)

	for _, arg := range table {
		lib.Challenge("Any", student.Any, correct.Any, arg.f, arg.a)
	}
}
