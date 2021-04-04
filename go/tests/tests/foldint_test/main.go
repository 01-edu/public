package main

import (
	"fmt"

	student "student"

	"github.com/01-edu/public/go/tests/lib"
)

func foldInt(f func(int, int) int, a []int, n int) {
	result := n
	for _, v := range a {
		result = f(result, v)
	}
	fmt.Println(result)
}

func main() {
	f := []func(int, int) int{
		func(accumulator, currentValue int) int {
			return accumulator + currentValue
		},
		func(accumulator, currentValue int) int {
			return accumulator - currentValue
		},
		func(accumulator, currentValue int) int {
			return currentValue * accumulator
		},
	}

	type node struct {
		a         []int
		functions []func(int, int) int
		n         int
	}
	argInt := []int{}
	table := []node{}

	for i := 0; i < 8; i++ {
		argInt = append(argInt, lib.MultRandIntBetween(0, 50)...)
		table = append(table, node{
			a:         argInt,
			functions: f,
			n:         lib.RandIntBetween(0, 60),
		})
	}

	table = append(table, node{
		a:         []int{1, 2, 3},
		functions: f,
		n:         93,
	})

	table = append(table, node{
		a:         []int{0},
		functions: f,
		n:         93,
	})

	for _, v := range table {
		for _, f := range v.functions {
			lib.Challenge("FoldInt", student.FoldInt, foldInt, f, v.a, v.n)
		}
	}
}
