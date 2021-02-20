package main

import (
	"math"

	student "student"

	"github.com/01-edu/public/go/tests/lib"
)

func iterativePower(nb int, power int) int {
	if power < 0 {
		return 0
	}
	result := math.Pow(float64(nb), float64(power))
	return int(result)
}

func main() {
	i := 0
	for i < 30 {
		nb := lib.RandIntBetween(-8, 8)
		power := lib.RandIntBetween(-10, 10)
		lib.Challenge("IterativePower", student.IterativePower, iterativePower, nb, power)
		i++
	}
	lib.Challenge("IterativePower", student.IterativePower, iterativePower, 0, 0)
	lib.Challenge("IterativePower", student.IterativePower, iterativePower, 0, 1)
}
