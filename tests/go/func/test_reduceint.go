package main

import (
	"../lib"
	"./correct"
	"./student"
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
		argInt = lib.MultRandIntBetween(0, 50)
		table = append(table, node{
			a:         argInt,
			functions: f,
		})
	}

	for _, v := range table {
		for _, f := range v.functions {
			lib.Challenge("ReduceInt", student.ReduceInt, correct.ReduceInt, f, v.a)
		}
	}
}
