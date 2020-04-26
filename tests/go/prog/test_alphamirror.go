package main

import "github.com/01-edu/z01"

func main() {
	args := [][]string{
		{""},
		{"One", "ring!"},
		{"testing spaces and #!*"},
		{"more", "than", "three", "arguments"},
		{"Upper anD LoWer cAsE"},
		{z01.RandWords()},
	}

	for _, v := range args {
		z01.ChallengeMain("alphamirror", v...)
	}
}
