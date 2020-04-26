package main

import (
	"github.com/01-edu/z01"

	correct "./correct"
	student "./student"
)

func main() {
	table := [][2]int{
		{50, 43},
		{13, 13},
		{10, 9},
		{5, 9},
	}
	for i := 0; i < 15; i++ {
		table = append(table, [2]int{
			z01.RandIntBetween(0, 30),
			z01.RandIntBetween(0, 100),
		})
	}
	for _, arg := range table {
		z01.Challenge("Volumechanger", student.Volumechanger, correct.Volumechanger, arg[0], arg[1])
	}
}
