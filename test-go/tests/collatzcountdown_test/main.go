package main

import (
	student "student"

	"github.com/01-edu/public/test-go/lib"
)

func collatzCountdown(start int) int {
	if start <= 0 {
		return -1
	}

	steps := 0

	for start != 1 {
		if start%2 == 0 {
			start = start / 2
		} else {
			start = 3*start + 1
		}
		steps++
	}

	return steps
}

func main() {
	args := append(lib.MultRandIntBetween(2, 20),
		lib.RandIntBetween(-6, 20),
		-5,
		0,
	)
	for _, v := range args {
		lib.Challenge("CollatzCountdown", student.CollatzCountdown, collatzCountdown, v)
	}
}
