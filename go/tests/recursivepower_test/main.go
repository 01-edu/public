package main

import (
	"math"

	student "student"

	"github.com/01-edu/public/go/tests/lib"
)

func recursivePower(nb, power int) int {
	if power < 0 {
		return 0
	}
	return int(math.Pow(float64(nb), float64(power)))
}

func main() {
	i := 0
	for i < 30 {
		nb := lib.RandIntBetween(-8, 8)
		power := lib.RandIntBetween(-10, 10)
		lib.Challenge("RecursivePower", student.RecursivePower, recursivePower, nb, power)
		i++
	}
	lib.Challenge("RecursivePower", student.RecursivePower, recursivePower, 0, 0)
	lib.Challenge("RecursivePower", student.RecursivePower, recursivePower, 0, 1)
}
