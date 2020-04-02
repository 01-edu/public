package main

import (
	"testing"

	"github.com/01-edu/z01"

	solutions "../../solutions" // This line is not necessary when testing an exercise with a program
)

func TestLcm(t *testing.T) {

	type node struct {
		first  int
		second int
	}

	table := []node{}

	table = append(table,
		node{50, 43},
		node{13, 13},
		node{10, 9},
		node{0, 9},
		node{1, 1},
	)

	// If we were to leave the table as it is, a student could just do a program with 4 ifs and get
	// "around" the goal of the exercise. We are now going to add 15 random tests using the z01 testing library

	for i := 0; i < 15; i++ {
		value := node{
			first:  z01.RandIntBetween(0, 1000),
			second: z01.RandIntBetween(0, 1000),
		}
		table = append(table, value)
	}

	for _, arg := range table {
		z01.Challenge(t, Lcm, solutions.Lcm, arg.first, arg.second)
	}
}
