package main

import (
	"github.com/01-edu/z01"

	solutions "./solutions"
	student "./student"
)

func main() {
	z01.Challenge("EightQueens", student.EightQueens, solutions.EightQueens)
}
