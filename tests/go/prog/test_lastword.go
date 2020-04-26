package main

import "github.com/01-edu/z01"

func main() {
	args := []string{
		"FOR PONY",
		"this ... is sparta, then again, maybe not",
		" ",
		" lorem,ipsum ",
	}

	args = append(args, z01.MultRandWords())

	for _, v := range args {
		z01.ChallengeMain("lastword", v)
	}

	z01.ChallengeMain("lastword", "a", "b")
	z01.ChallengeMain("lastword")
}
