package main

import (
	"./student"
	"github.com/01-edu/public/go/lib"
	"github.com/01-edu/public/go/lib/is"
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
	functions := []func(string) bool{is.Digit, is.Lower, is.Upper}

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
			f:   is.Digit,
			arr: lib.MultRandDigit(),
		}
		table = append(table, val)
	}

	for i := 0; i < 5; i++ {
		val := node{
			f:   is.Lower,
			arr: lib.MultRandLower(),
		}
		table = append(table, val)
	}
	for i := 0; i < 5; i++ {
		val := node{
			f:   is.Upper,
			arr: lib.MultRandUpper(),
		}
		table = append(table, val)
	}

	table = append(table,
		node{
			f:   is.Digit,
			arr: []string{"Hello", "how", "are", "you"},
		},
		node{
			f:   is.Digit,
			arr: []string{"This", "is", "4", "you"},
		},
	)

	for _, arg := range table {
		lib.Challenge("CountIf", student.CountIf, countIf, arg.f, arg.arr)
	}
}
