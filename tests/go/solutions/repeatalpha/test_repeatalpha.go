package main

import "github.com/01-edu/z01"

func main() {
	args := []string{
		"",
		"Hello",
		"World",
		"Home",
		"Theorem",
		"Choumi is the best cat",
		"abracadaba 01!",
		"abc",
		"MaTheMatiCs",
	}

	args = append(args, z01.MultRandAlnum()...)

	for _, v := range args {
		z01.ChallengeMain("repeatalpha", v)
	}
	z01.ChallengeMain("repeatalpha")
	z01.ChallengeMain("repeatalpha", "", "")
}
