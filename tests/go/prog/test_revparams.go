package main

import "../lib"

func main() {
	lib.ChallengeMain("revparams", "choumi", "is", "the", "best", "cat")
	lib.ChallengeMain("revparams", lib.MultRandWords()...)
}
