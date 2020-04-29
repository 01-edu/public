package main

import (
	"../lib"
	"./correct"
	"./student"
)

func main() {
	args := []int{lib.RandIntBetween(-6, 20)}
	args = append(args, -5, 0)
	for i := 0; i < 20; i++ {
		args = append(args, lib.RandIntBetween(2, 20))
	}

	for _, v := range args {
		lib.Challenge("CollatzCountdown", student.CollatzCountdown, correct.CollatzCountdown, v)
	}
}
