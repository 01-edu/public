package main

import (
	"github.com/01-edu/public/go/tests/lib"
)

func main() {
	args := [][]string{
		{"hello", "you"},
		{"   only  it's harder   "},
		{"you   see   it's   easy   to   display    the     same  thing"},
	}

	args = append(args, lib.MultRandWords())

	for _, v := range args {
		lib.ChallengeMain("expandstr", v...)
	}
	lib.ChallengeMain("expandstr")
}
