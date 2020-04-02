package main

import (
	"testing"

	solutions "../../solutions"
	"github.com/01-edu/z01"
)

func TestReduceInt(t *testing.T) {
	f := []func(int, int) int{Add, Sub, Mul}
	argInt := []int{}

	type node struct {
		arr       []int
		functions []func(int, int) int
	}

	table := []node{}
	for i := 0; i < 4; i++ {
		argInt = z01.MultRandIntBetween(0, 50)
		val := node{
			arr:       argInt,
			functions: f,
		}
		table = append(table, val)
	}

	for _, v := range table {
		for _, f := range v.functions {
			z01.Challenge(t, ReduceInt, solutions.ReduceInt, f, v.arr)
		}
	}
}

func Add(accumulator, currentValue int) int {
	return accumulator + currentValue
}

func Sub(accumulator, currentValue int) int {
	return accumulator - currentValue
}

func Mul(accumulator, currentValue int) int {
	return currentValue * accumulator
}
