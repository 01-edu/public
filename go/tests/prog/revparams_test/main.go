package main

import (
	"github.com/01-edu/public/go/tests/lib"
)

func main() {
	lib.ChallengeMain("revparams", "choumi", "is", "the", "best", "cat")
	lib.ChallengeMain("revparams", lib.MultRandWords()...)
}
