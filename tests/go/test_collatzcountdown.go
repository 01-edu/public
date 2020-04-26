package main

import (
	"github.com/01-edu/z01"

	solutions "./solutions"
	student "./student"
)

func main() {
	args := []int{z01.RandIntBetween(-6, 20)}
	args = append(args, -5, 0)
	for i := 0; i < 20; i++ {
		args = append(args, z01.RandIntBetween(2, 20))
	}

	for _, v := range args {
		z01.Challenge("CollatzCountdown", student.CollatzCountdown, solutions.CollatzCountdown, v)
	}
}
