package main

import "github.com/01-edu/z01"

func main() {
	args := [][]string{
		{"hello", "you"},
		{"   only  it's harder   "},
		{"you   see   it's   easy   to   display    the     same  thing"},
	}

	args = append(args, z01.MultRandWords())

	for _, v := range args {
		z01.ChallengeMain("expandstr", v...)
	}
	z01.ChallengeMain("expandstr")
}
