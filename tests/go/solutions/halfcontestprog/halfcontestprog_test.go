package main

import (
	"testing"

	"github.com/01-edu/z01"

	solutions "../../solutions" // This line is not necessary when testing an exercise with a program
)

func TestHalf_contest(t *testing.T) {
	type node struct {
		h1 int
		m1 int
		h2 int
		m2 int
	}
	table := []node{}
	// Initial filling of that array with the values I see in the examples of the subject

	table = append(table,
		node{11, 44, 21, 59},
		node{1, 12, 1, 14},
		node{5, 50, 6, 51},
		node{14, 35, 18, 55},
	)

	// If we were to leave the table as it is, a student could just do a program with 4 ifs and get
	// "around" the goal of the exercise. We are now going to add 15 random tests using the z01 testing library

	for i := 0; i < 20; i++ {
		value := node{
			h1: z01.RandIntBetween(0, 10),
			m1: z01.RandIntBetween(0, 59),
			h2: z01.RandIntBetween(11, 23),
			m2: z01.RandIntBetween(0, 59),
		}

		//Once the random node created, this iteration is added to the earlier declared table
		//along with the 4 specific examples taken from the examples of the readme.
		table = append(table, value)
	}

	//Once the random node created, this iteration is added to the earlier declared table
	//along with the 4 specific examples taken from the examples of the readme.
	//The table with 4 specific exercises and 15 randoms is now ready to be "challenged"
	//Because the exercise asks for a function we are now using the Challenge function (this function would
	// be the ChallengeMainExam function)

	for _, arg := range table {
		z01.Challenge(t, Halfcontest, solutions.Halfcontest, arg.h1, arg.m1, arg.h2, arg.m2)
	}

}
