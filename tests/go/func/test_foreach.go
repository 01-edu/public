package main

import (
	"fmt"

	"../lib"
	"./correct"
	"./student"
)

func add0(i int) {
	fmt.Println(i)
}
func add1(i int) {
	fmt.Println(i + 1)
}
func add2(i int) {
	fmt.Println(i + 2)
}
func add3(i int) {
	fmt.Println(i + 3)
}

func main() {
	functions := []func(int){add0, add1, add2, add3}

	type node struct {
		f func(int)
		a []int
	}

	table := []node{}

	// 15 random slice of random ints with a random function from selection
	for i := 0; i < 15; i++ {
		function := functions[lib.RandIntBetween(0, len(functions)-1)]
		table = append(table, node{
			f: function,
			a: lib.MultRandIntBetween(-1000000, 1000000),
		})
	}

	table = append(table, node{
		f: add0,
		a: []int{1, 2, 3, 4, 5, 6},
	})

	for _, arg := range table {
		lib.Challenge("ForEach", student.ForEach, correct.ForEach, arg.f, arg.a)
	}
}
