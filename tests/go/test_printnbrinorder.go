package main

import (
	"github.com/01-edu/z01"

	solutions "./solutions"
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
		z01.Challenge("PrintNbrInOrder", student.PrintNbrInOrder, solutions.PrintNbrInOrder, arg)
	}
}
