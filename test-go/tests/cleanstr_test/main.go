package main

import (
	"github.com/01-edu/public/test-go/lib"
)

func main() {
	args := []string{
		"you see it's easy to display the same thing",
		" only  it's   harder  ",
		"how   funny",
		"",
		lib.RandSpace(),
	}

	for _, v := range args {
		lib.ChallengeMain("cleanstr", v)
	}
	lib.ChallengeMain("cleanstr", "this is   not", "happening")
}
