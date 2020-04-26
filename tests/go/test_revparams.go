package main

import "github.com/01-edu/z01"

func main() {
	z01.ChallengeMain("revparams", "choumi", "is", "the", "best", "cat")
	z01.ChallengeMain("revparams", z01.MultRandWords()...)
}
