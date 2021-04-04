package main

import (
	"github.com/01-edu/public/test-go/lib"
)

func main() {
	lib.ChallengeMain("revparams", "choumi", "is", "the", "best", "cat")
	lib.ChallengeMain("revparams", lib.MultRandWords()...)
}
