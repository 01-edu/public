package main

import (
	student "student"

	"github.com/01-edu/public/go/tests/lib"
)

func min3(a, b, c int) int {
	if a <= b && a <= c {
		return a
	}
	if b <= a && b <= c {
		return b
	}
	return c
}

func max3(a, b, c int) int {
	if a >= b && a >= c {
		return a
	}
	if b >= a && b >= c {
		return b
	}
	return c
}

func min2(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func sweetproblem(a, b, c int) int {
	if a > b {
		f := a
		a = b
		b = f
	}
	if a > c {
		f := a
		a = c
		c = f
	}
	if b > c {
		f := b
		b = c
		c = f
	}
	ans := a
	if c-b >= a {
		c -= a
	} else {
		a -= c - b
		half := a / 2
		c -= half
		b -= a - half
	}
	ans += min2(b, c)
	return ans
}

func main() {
	type node struct {
		red   int
		green int
		blue  int
	}

	table := []node{
		{50, 43, 20},
		{10, 0, 0},
		{0, 10, 0},
		{0, 0, 10},
		{10, 1, 0},
		{0, 10, 1},
		{1, 0, 10},
		{10, 2, 0},
		{2, 10, 0},
		{0, 2, 10},
		{13, 13, 0},
		{10, 9, 0},
		{5, 9, 2},
	}

	for i := 0; i < 15; i++ {
		table = append(table, node{
			red:   lib.RandIntBetween(0, 30),
			green: lib.RandIntBetween(0, 30),
			blue:  lib.RandIntBetween(0, 30),
		})
	}
	for _, arg := range table {
		lib.Challenge("SweetProblem", student.Sweetproblem, sweetproblem, arg.red, arg.green, arg.blue)
	}
}
