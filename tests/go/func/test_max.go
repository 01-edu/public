package main

import (
	"../lib"
	"./correct"
	"./student"
)

func main() {
	args := []int{lib.RandInt()}
	limit := lib.RandIntBetween(20, 50)
	for i := 0; i < limit; i++ {
		args = append(args, lib.RandInt())
	}

	lib.Challenge("Max", student.Max, correct.Max, args)
}
