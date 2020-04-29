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
		lib.Challenge("IterativePower", student.IterativePower, correct.IterativePower, nb, power)
		i++
	}
	lib.Challenge("IterativePower", student.IterativePower, correct.IterativePower, 0, 0)
	lib.Challenge("IterativePower", student.IterativePower, correct.IterativePower, 0, 1)
}
