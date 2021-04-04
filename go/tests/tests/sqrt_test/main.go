package main

import (
	"math"

	student "student"

	"github.com/01-edu/public/go/tests/lib"
)

func sqrt(value int) int {
	sr := math.Sqrt(float64(value))
	if math.Mod(sr, 1) == 0 {
		return int(sr)
	}
	return 0
}

func main() {
	table := append(lib.MultRandIntBetween(-1000000, 1000000),
		0,
		1,
		2,
		3,
		4,
		5,
		6,
		7,
		8,
		9,
		10,
		11,
		12,
		100,
	)
	for _, arg := range table {
		lib.Challenge("Sqrt", student.Sqrt, sqrt, arg)
	}
}
