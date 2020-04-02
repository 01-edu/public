package main

import (
	"testing"

	"github.com/01-edu/z01"
)

func TestRoString(t *testing.T) {
	corrArgs := []string{"abc   ",
		"Let there     be light",
		"     AkjhZ zLKIJz , 23y",
		"",
	}

	// some random valid parameter
	for i := 0; i < 5; i++ {
		corrArgs = append(corrArgs, z01.RandWords())
	}

	for _, v := range corrArgs {
		z01.ChallengeMainExam(t, v)
	}

	//without parameter
	z01.ChallengeMainExam(t)

	//with more than one parameter
	z01.ChallengeMainExam(t, "this", "is")
	z01.ChallengeMainExam(t, "not", "good", "for  you")
}
