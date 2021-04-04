package main

import (
	"github.com/01-edu/public/go/tests/lib"
)

func main() {
	args := []string{
		"FOR PONY",
		"this ... is sparta, then again, maybe not",
		" ",
		" lorem,ipsum ",
	}

	args = append(args, lib.MultRandWords()...)

	for _, v := range args {
		lib.ChallengeMain("lastword", v)
	}

	lib.ChallengeMain("lastword", "a", "b")
	lib.ChallengeMain("lastword")
}
