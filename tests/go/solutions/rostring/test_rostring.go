package main

import (
	"github.com/01-edu/z01"
)

func main() {
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
		z01.ChallengeMain("rostring", v)
	}

	//without parameter
	z01.ChallengeMain("rostring")

	//with more than one parameter
	z01.ChallengeMain("rostring", "this", "is")
	z01.ChallengeMain("rostring", "not", "good", "for  you")
}
