package main

import (
	"github.com/01-edu/z01"

	solutions "./solutions"
	student "./student"
)

func main() {
	table := append(
		z01.MultRandASCII(),
		"Hello! How are you?",
	)
	for _, arg := range table {
		z01.Challenge("ToLower", student.ToLower, solutions.ToLower, arg)
	}
}
