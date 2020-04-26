package main

import (
	"github.com/01-edu/z01"

	correct "./correct"
	student "./student"
)

func main() {
	functions := []func(string) bool{correct.IsNumeric, correct.IsLower, correct.IsUpper}

	type node struct {
		f func(string) bool
		a []string
	}

	table := []node{}

	for i := 0; i < 5; i++ {
		function := functions[z01.RandIntBetween(0, len(functions)-1)]
		table = append(table, node{
			f: function,
			a: z01.MultRandWords(),
		})
	}
	for i := 0; i < 5; i++ {
		table = append(table, node{
			f: correct.IsNumeric,
			a: z01.MultRandDigit(),
		})
	}

	for i := 0; i < 5; i++ {
		table = append(table, node{
			f: correct.IsLower,
			a: z01.MultRandLower(),
		})
	}
	for i := 0; i < 5; i++ {
		table = append(table, node{
			f: correct.IsUpper,
			a: z01.MultRandUpper(),
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
		z01.Challenge("Any", student.Any, correct.Any, arg.f, arg.a)
	}
}
