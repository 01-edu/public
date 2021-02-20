package main

import (
	"github.com/01-edu/public/go/tests/lib"
)

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

	args = append(args, lib.MultRandAlnum()...)

	for _, v := range args {
		lib.ChallengeMain("repeatalpha", v)
	}
	lib.ChallengeMain("repeatalpha")
	lib.ChallengeMain("repeatalpha", "", "")
}
