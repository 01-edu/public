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
	for i := 0; i < 25; i++ {
		table = append(table, z01.MultRandIntBetween(1, 877))
	}
	for _, arg := range table {
		z01.Challenge("ReachableNumber", student.ReachableNumber, correct.ReachableNumber, arg)
	}
}
