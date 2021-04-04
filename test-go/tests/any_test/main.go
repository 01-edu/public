package main

import (
	student "student"

	"github.com/01-edu/public/test-go/lib"
	"github.com/01-edu/public/test-go/lib/is"
)

func any(f func(string) bool, a []string) bool {
	for _, el := range a {
		if f(el) {
			return true
		}
	}
	return false
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
		table = append(table, node{
			f: function,
			a: lib.MultRandWords(),
		})
	}
	for i := 0; i < 5; i++ {
		table = append(table, node{
			f: is.Digit,
			a: lib.MultRandDigit(),
		})
	}

	for i := 0; i < 5; i++ {
		table = append(table, node{
			f: is.Lower,
			a: lib.MultRandLower(),
		})
	}
	for i := 0; i < 5; i++ {
		table = append(table, node{
			f: is.Upper,
			a: lib.MultRandUpper(),
		})
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
		lib.Challenge("Any", student.Any, any, arg.f, arg.a)
	}
}
