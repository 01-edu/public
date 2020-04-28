package main

import "../lib"

func main() {
	args := []string{
		"abc   ",
		"Let there     be light",
		"     AkjhZ zLKIJz , 23y",
		"",
	}
	args = append(args, lib.MultRandWords()...)

	for _, arg := range args {
		lib.ChallengeMain("rostring", arg)
	}
	lib.ChallengeMain("rostring")
	lib.ChallengeMain("rostring", "this", "is")
	lib.ChallengeMain("rostring", "not", "good", "for  you")
}
