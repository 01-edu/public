package main

import (
	"../lib"
	"./correct"
	"./student"
)

func main() {
	i := 0
	for i < 30 {
		nb := lib.RandIntBetween(-8, 8)
		power := lib.RandIntBetween(-10, 10)
		lib.Challenge("RecursivePower", student.RecursivePower, correct.RecursivePower, nb, power)
		i++
	}
	lib.Challenge("RecursivePower", student.RecursivePower, correct.RecursivePower, 0, 0)
	lib.Challenge("RecursivePower", student.RecursivePower, correct.RecursivePower, 0, 1)
}
