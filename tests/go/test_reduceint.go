package main

import (
	"github.com/01-edu/z01"

	correct "./correct"
	student "./student"
)

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
	argInt := []int{}

	type node struct {
		a         []int
		functions []func(int, int) int
	}

	table := []node{}
	for i := 0; i < 4; i++ {
		argInt = z01.MultRandIntBetween(0, 50)
		table = append(table, node{
			a:         argInt,
			functions: f,
		})
	}

	for _, v := range table {
		for _, f := range v.functions {
			z01.Challenge("ReduceInt", student.ReduceInt, correct.ReduceInt, f, v.a)
		}
	}
}
