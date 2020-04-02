package main

import (
	"testing"

	solutions "../../solutions"
	"github.com/01-edu/z01"
)

func TestFoldInt(t *testing.T) {
	f := []func(int, int) int{Add, Sub, Mul}

	type node struct {
		arr       []int
		functions []func(int, int) int
		n         int
	}
	argInt := []int{}
	table := []node{}

	for i := 0; i < 8; i++ {
		argInt = append(argInt, z01.MultRandIntBetween(0, 50)...)
		val := node{
			arr:       argInt,
			functions: f,
			n:         z01.RandIntBetween(0, 60),
		}
		table = append(table, val)
	}

	table = append(table, node{
		arr:       []int{1, 2, 3},
		functions: f,
		n:         93,
	})

	table = append(table, node{
		arr:       []int{0},
		functions: f,
		n:         93,
	})

	for _, v := range table {
		for _, f := range v.functions {
			z01.Challenge(t, FoldInt, solutions.FoldInt, f, v.arr, v.n)
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
