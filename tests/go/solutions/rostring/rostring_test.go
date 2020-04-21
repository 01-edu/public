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
		z01.ChallengeMain(v)
	}

	//without parameter
	z01.ChallengeMain()

	//with more than one parameter
	z01.ChallengeMain("this", "is")
	z01.ChallengeMain("not", "good", "for  you")
}
