package main

import (
	"../common"
	"./student"
	"github.com/01-edu/public/go/lib"
)

func any(f func(string) bool, arr []string) bool {
	for _, el := range arr {
		if f(el) {
			return true
		}
	}

	return false
}

func main() {
	functions := []func(string) bool{common.IsDigit, common.IsLower, common.IsUpper}

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
			f: common.IsDigit,
			a: lib.MultRandDigit(),
		})
	}

	for i := 0; i < 5; i++ {
		table = append(table, node{
			f: common.IsLower,
			a: lib.MultRandLower(),
		})
	}
	for i := 0; i < 5; i++ {
		table = append(table, node{
			f: common.IsUpper,
			a: lib.MultRandUpper(),
		})
	}

	table = append(table,
		node{
			f: common.IsDigit,
			a: []string{"Hello", "how", "are", "you"},
		},
		node{
			f: common.IsDigit,
			a: []string{"This", "is", "4", "you"},
		},
	)

	for _, arg := range table {
		lib.Challenge("Any", student.Any, any, arg.f, arg.a)
	}
}
