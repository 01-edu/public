package main

import (
	"testing"

	"github.com/01-edu/z01"

	solutions "../../solutions" // This line is not necessary when testing an exercise with a program
)

func TestSweerproblem(t *testing.T) {
	type node struct {
		red   int
		green int
		blue  int
	}

	table := []node{}
	// Initial filling of that array with the values I see in the examples of the subject

	table = append(table,
		node{50, 43, 20},
		node{10, 0, 0},
		node{0, 10, 0},
		node{0, 0, 10},
		node{10, 1, 0},
		node{0, 10, 1},
		node{1, 0, 10},
		node{10, 2, 0},
		node{2, 10, 0},
		node{0, 2, 10},
		node{13, 13, 0},
		node{10, 9, 0},
		node{5, 9, 2},
	)

	// If we were to leave the table as it is, a student could just do a program with 4 ifs and get
	// "around" the goal of the exercise. We are now going to add 15 random tests using the z01 testing library

	for i := 0; i < 15; i++ {
		value := node{
			red:   z01.RandIntBetween(0, 30),
			green: z01.RandIntBetween(0, 30),
			blue:  z01.RandIntBetween(0, 30),
		}

		//Once the random node created, this iteration is added to the earlier declared table
		//along with the 4 specific examples taken from the examples of the readme.
		table = append(table, value)

	}

	//The table with 4 specific exercises and 15 randoms is now ready to be "challenged"
	//Because the exercise asks for a function we are now using the Challenge function (this function would
	// be the ChallengeMainExam function)

	for _, arg := range table {
		z01.Challenge(t, Sweetproblem, solutions.Sweetproblem, arg.red, arg.green, arg.blue)
	}
}
