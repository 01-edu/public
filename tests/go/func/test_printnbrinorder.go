package main

import (
	"github.com/01-edu/z01"

	correct "./correct"
	student "./student"
)

func main() {
	table := append(
		z01.MultRandIntBetween(0, z01.MaxInt),
		z01.MaxInt,
		321,
		321321,
		0,
	)
	for _, arg := range table {
		z01.Challenge("PrintNbrInOrder", student.PrintNbrInOrder, correct.PrintNbrInOrder, arg)
	}
}
