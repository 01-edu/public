package main

import (
	"testing"

	"github.com/01-edu/z01"

	solutions "../../solutions" // This line is not necessary when testing an exercise with a program
)

func TestReachableNumber(t *testing.T) {
	type node struct {
		n int
	}

	table := []node{}
	// Initial filling of that array with the values I see in the examples of the subject

	table = append(table,
		node{20},
		node{1},
		node{9},
		node{2},
	)

	// If we were to leave the table as it is, a student could just do a program with 4 ifs and get
	// "around" the goal of the exercise. We are now going to add 15 random tests using the z01 testing library

	for i := 0; i < 25; i++ {
		value := node{
			n: z01.RandIntBetween(1, 877),
		}

		//Once the random node created, this iteration is added to the earlier declared table
		//along with the 4 specific examples taken from the examples of the readme.
		table = append(table, value)

	}

	//The table with 4 specific exercises and 15 randoms is now ready to be "challenged"
	//Because the exercise asks for a function we are now using the Challenge function (this function would
	// be the ChallengeMainExam function)

	for _, arg := range table {
		z01.Challenge(t, Reachablenumber, solutions.Reachablenumber, arg.n)
	}
}
