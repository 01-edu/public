package main

import (
	"github.com/01-edu/z01"

	correct "./correct"
	student "./student"
)

func main() {
	table := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, arg := range table {
		z01.Challenge("PrintCombN", student.PrintCombN, correct.PrintCombN, arg)
	}
}
