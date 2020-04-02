package main

import (
	"testing"

	"github.com/01-edu/z01"

	solutions "../../solutions" // This line is not necessary when testing an exercise with a program
)

func TestPriorprime(t *testing.T) {

	// Declaration of the node that is going to take the group of arguments that are going to
	// inputed during each iteration of a Challenge between the student and the staff solution.
	// (note: a node is not always necessary but in this case it makes the writing of the test easier)

	type node struct {
		first  int
	}

	// Declaration of  an empty array of type node{}
	// note that in this case this is the easiest type of table to declare
	// but a table can be of any other relevant type, (for example []string{}, []int{} if it
	// were a single string tested or a single int)

	table := []node{}
	// Initial filling of that array with the values I see in the examples of the subject

	table = append(table,
		node{50},
		node{13},
		node{10},
		node{0},
		node{1},		
		node{2},		
	)

	// If we were to leave the table as it is, a student could just do a program with 4 ifs and get
	// "around" the goal of the exercise. We are now going to add 15 random tests using the z01 testing library

	for i := 0; i < 15; i++ {
		value := node{
			first:  z01.RandIntBetween(0, 1000),
			//this z01.RandIntBetween function allows the randomization of
			//the int for each value in a desired range.
			//Note that they are many others of those functions for other types of data
			//Do not hesitate to have a look at all of them https://github.com/01-edu/z01
		}

		//Once the random node created, this iteration is added to the earlier declared table
		//along with the 4 specific examples taken from the examples of the readme.
		table = append(table, value)

	}

	//The table with 4 specific exercises and 15 randoms is now ready to be "challenged"
	//Because the exercise asks for a function we are now using the Challenge function (this function would
	// be the ChallengeMainExam function)

	for _, arg := range table {
		z01.Challenge(t, Priorprime, solutions.Priorprime, arg.first)
	}

	// the z01.Challenge function is here applied to each argument of the table. It musts contains:
	// 1) first, the t argument from the T structure imported from the package "testing"
	//
	// 2) second, the function from the student, in this case Nauuo
	//(this disapears in the ChallengeMainExam function)
	// 3) third, the function from the staff, in this case solutions.Nauuo
	//(this disapears as well in the ChallengeMainExam function)
	// 4) all the arguments to be tested, in this case it is the plus, minus and rand from each structure,
	// notice that they are accessed with arg. (the arg notation comes from the way it was name in the
	// range loop over the table)

	// Now that this is done. re-read the quickReadme (the test your test recap) and apply all the commands
	// and intructions. We strongly advise to check that your error messages matches your subject.
	// and that you ask a colleague to double check.

	//FINAL STEP:
	// When both are satisfied with the coherence between the subject and its tests. The code can be commited
	// and redeployed by the team-01.
	// We then advised the staff team to test the new exercise invidually with their current build of the exam

}
