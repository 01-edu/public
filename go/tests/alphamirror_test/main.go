package main

import (
	"github.com/01-edu/public/go/tests/lib"
)

func main() {
	args := [][]string{
		{""},
		{"One", "ring!"},
		{"testing spaces and #!*"},
		{"more", "than", "three", "arguments"},
		{"Upper anD LoWer cAsE"},
		{lib.RandWords()},
	}

	for _, v := range args {
		lib.ChallengeMain("alphamirror", v...)
	}
}
