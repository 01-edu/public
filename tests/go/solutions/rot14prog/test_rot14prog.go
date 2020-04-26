package main

import "github.com/01-edu/z01"

func main() {
	table := append(z01.MultRandWords(),
		z01.MultRandWords()...,
	)

	for _, arg := range table {
		z01.ChallengeMain("rot14prog", arg)
	}
	z01.ChallengeMain("rot14prog", "", "something", "something1")
}
