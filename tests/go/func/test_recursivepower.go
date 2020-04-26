package main

import (
	"github.com/01-edu/z01"

	correct "./correct"
	student "./student"
)

func main() {
	i := 0
	for i < 30 {
		nb := z01.RandIntBetween(-8, 8)
		power := z01.RandIntBetween(-10, 10)
		z01.Challenge("RecursivePower", student.RecursivePower, correct.RecursivePower, nb, power)
		i++
	}
	z01.Challenge("RecursivePower", student.RecursivePower, correct.RecursivePower, 0, 0)
	z01.Challenge("RecursivePower", student.RecursivePower, correct.RecursivePower, 0, 1)
}
