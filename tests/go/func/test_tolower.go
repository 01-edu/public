package main

import (
	"github.com/01-edu/z01"

	correct "./correct"
	student "./student"
)

func main() {
	table := append(
		z01.MultRandASCII(),
		"Hello! How are you?",
	)
	for _, arg := range table {
		z01.Challenge("ToLower", student.ToLower, correct.ToLower, arg)
	}
}
