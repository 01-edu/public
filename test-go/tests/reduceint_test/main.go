package main

import (
	"fmt"

	student "student"

	"github.com/01-edu/public/test-go/lib"
)

func reduceInt(a []int, f func(int, int) int) {
	acc := a[0]
	for i := 1; i < len(a); i++ {
		acc = f(acc, a[i])
	}
	fmt.Println(acc)
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
			lib.Challenge("ReduceInt", student.ReduceInt, reduceInt, v.a, f)
		}
	}
}
