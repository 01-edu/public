package main

import "github.com/01-edu/z01"

func main() {
	args := []string{
		"abc   ",
		"Let there     be light",
		"     AkjhZ zLKIJz , 23y",
		"",
	}
	args = append(args, z01.MultRandWords()...)

	for _, arg := range args {
		z01.ChallengeMain("rostring", arg)
	}
	z01.ChallengeMain("rostring")
	z01.ChallengeMain("rostring", "this", "is")
	z01.ChallengeMain("rostring", "not", "good", "for  you")
}
