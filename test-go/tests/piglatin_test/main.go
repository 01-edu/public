package main

import (
	"github.com/01-edu/public/test-go/lib"
)

func main() {
	args := [][]string{
		{"", "pig", "is", "crunch", "crnch"},
		{"something", "else"},
	}

	for i := 0; i < 4; i++ {
		args = append(args, lib.MultRandBasic())
	}

	for _, v := range args {
		lib.ChallengeMain("piglatin", v...)
	}
}
