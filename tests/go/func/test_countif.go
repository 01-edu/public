package main

import (
	"github.com/01-edu/z01"

	correct "./correct"
	student "./student"
)

func main() {
	functions := []func(string) bool{correct.IsNumeric, correct.IsLower, correct.IsUpper}

	type node struct {
		f   func(string) bool
		arr []string
	}

	table := []node{}

	for i := 0; i < 5; i++ {
		function := functions[z01.RandIntBetween(0, len(functions)-1)]
		val := node{
			f:   function,
			arr: z01.MultRandWords(),
		}
		table = append(table, val)
	}
	for i := 0; i < 5; i++ {
		val := node{
			f:   correct.IsNumeric,
			arr: z01.MultRandDigit(),
		}
		table = append(table, val)
	}

	for i := 0; i < 5; i++ {
		val := node{
			f:   correct.IsLower,
			arr: z01.MultRandLower(),
		}
		table = append(table, val)
	}
	for i := 0; i < 5; i++ {
		val := node{
			f:   correct.IsUpper,
			arr: z01.MultRandUpper(),
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
		z01.Challenge("CountIf", student.CountIf, correct.CountIf, arg.f, arg.arr)
	}
}
