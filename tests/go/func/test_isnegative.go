package main

import (
	"github.com/01-edu/z01"

	correct "./correct"
	student "./student"
)

func main() {
	table := append(
		z01.MultRandInt(),
		z01.MinInt,
		z01.MaxInt,
		0,
	)
	for _, arg := range table {
		z01.Challenge("IsNegative", student.IsNegative, correct.IsNegative, arg)
	}
}
