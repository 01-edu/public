package main

import (
	"github.com/01-edu/z01"

	correct "./correct"
	student "./student"
)

func main() {
	table := []int{
		20,
		1,
		9,
		2,
	}
	table = append(table, z01.MultRandIntBetween(1, 1500)...)

	for _, arg := range table {
		z01.Challenge("InterestingNumber", student.InterestingNumber, correct.InterestingNumber, arg)
	}
}
