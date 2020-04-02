package main

import (
	"testing"

	"github.com/01-edu/z01"

	solutions "../../solutions" // This line is not necessary when testing an exercise with a program
)

func TestVolumechanger(t *testing.T) {
	type node struct {
		init int
		fin  int
	}
	table := []node{}
	// Initial filling of that array with the values I see in the examples of the subject

	table = append(table,
		node{50, 43},
		node{13, 13},
		node{10, 9},
		node{5, 9},
	)

	// If we were to leave the table as it is, a student could just do a program with 4 ifs and get
	// "around" the goal of the exercise. We are now going to add 15 random tests using the z01 testing library

	for i := 0; i < 15; i++ {
		value := node{
			init: z01.RandIntBetween(0, 30),
			fin:  z01.RandIntBetween(0, 100),
		}

		//Once the random node created, this iteration is added to the earlier declared table
		//along with the 4 specific examples taken from the examples of the readme.
		table = append(table, value)

	}

	//The table with 4 specific exercises and 15 randoms is now ready to be "challenged"
	//Because the exercise asks for a function we are now using the Challenge function (this function would
	// be the ChallengeMainExam function)

	for _, arg := range table {
		z01.Challenge(t, Volumechanger, solutions.Volumechanger, arg.init, arg.fin)
	}
}
