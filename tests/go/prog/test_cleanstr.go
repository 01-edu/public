package main

import "github.com/01-edu/z01"

func main() {
	args := []string{
		"you see it's easy to display the same thing",
		" only  it's   harder  ",
		"how   funny",
		"",
		z01.RandSpace(),
	}

	for _, v := range args {
		z01.ChallengeMain("cleanstr", v)
	}
	z01.ChallengeMain("cleanstr", "this is   not", "happening")
}
