package main

import (
	"github.com/01-edu/z01"
)

func main() {
	args := []string{"Hello",
		"World",
		"Home",
		"Theorem",
		"Choumi is the best cat",
		"abracadaba 01!",
		"abc",
		"MaTheMatiCs"}

	for i := 0; i < 5; i++ {
		args = append(args, z01.RandAlnum())
	}

	for _, v := range args {
		z01.ChallengeMain(v)
	}
}
