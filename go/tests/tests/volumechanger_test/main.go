package main

import (
	student "student"

	"github.com/01-edu/public/go/tests/lib"
)

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func volumechanger(a, b int) int {
	return abs(a-b)/5 + abs(a-b)%5/2 + abs(a-b)%5%2
}

func main() {
	table := [][2]int{
		{50, 43},
		{13, 13},
		{10, 9},
		{5, 9},
	}
	for i := 0; i < 15; i++ {
		table = append(table, [2]int{
			lib.RandIntBetween(0, 30),
			lib.RandIntBetween(0, 100),
		})
	}
	for _, arg := range table {
		lib.Challenge("Volumechanger", student.Volumechanger, volumechanger, arg[0], arg[1])
	}
}
