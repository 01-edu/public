package main

import (
	"strings"

	"github.com/01-edu/z01"

	solutions "./solutions"
	student "./student"
)

func main() {
	table := []string{}
	//30 random slice of strings

	for i := 0; i < 30; i++ {
		val := strings.Join(z01.MultRandASCII(), " ")
		table = append(table, val)
	}

	table = append(table,
		"Hello how are you?")

	for _, arg := range table {
		z01.Challenge("SplitWhiteSpaces", student.SplitWhiteSpaces, solutions.SplitWhiteSpaces, arg)
	}
}
