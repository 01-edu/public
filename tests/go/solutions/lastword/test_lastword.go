package main

import (
	"github.com/01-edu/z01"
)

func main() {
	args := []string{
		"FOR PONY",
		"this ... is sparta, then again, maybe not",
		" ",
		" lorem,ipsum ",
	}

	for i := 0; i < 5; i++ {
		args = append(args, z01.RandWords())
	}

	for _, v := range args {
		z01.ChallengeMain("lastword", v)
	}

	z01.ChallengeMain("lastword", "a", "b")
	z01.ChallengeMain("lastword")
}
