package main

import (
	student "student"

	"github.com/01-edu/public/go/tests/lib"
	"github.com/01-edu/public/go/tests/lib/is"
)

func countIf(f func(string) bool, a []string) int {
	counter := 0
	for _, el := range a {
		if f(el) {
			counter++
		}
	}

	return counter
}

func main() {
	functions := []func(string) bool{is.Digit, is.Lower, is.Upper}

	type node struct {
		f func(string) bool
		a []string
	}

	table := []node{}

	for i := 0; i < 5; i++ {
		function := functions[lib.RandIntBetween(0, len(functions)-1)]
		val := node{
			f: function,
			a: lib.MultRandWords(),
		}
		table = append(table, val)
	}
	for i := 0; i < 5; i++ {
		val := node{
			f: is.Digit,
			a: lib.MultRandDigit(),
		}
		table = append(table, val)
	}

	for i := 0; i < 5; i++ {
		val := node{
			f: is.Lower,
			a: lib.MultRandLower(),
		}
		table = append(table, val)
	}
	for i := 0; i < 5; i++ {
		val := node{
			f: is.Upper,
			a: lib.MultRandUpper(),
		}
		table = append(table, val)
	}

	table = append(table,
		node{
			f: is.Digit,
			a: []string{"Hello", "how", "are", "you"},
		},
		node{
			f: is.Digit,
			a: []string{"This", "is", "4", "you"},
		},
	)

	for _, arg := range table {
		lib.Challenge("CountIf", student.CountIf, countIf, arg.f, arg.a)
	}
}
