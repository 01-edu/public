package main

import "github.com/01-edu/z01"

func main() {
	args := [][]string{
		{"", "pig", "is", "crunch", "crnch"},
		{"something", "else"},
	}

	for i := 0; i < 4; i++ {
		args = append(args, z01.MultRandBasic())
	}

	for _, v := range args {
		z01.ChallengeMain("piglatin", v...)
	}
}
