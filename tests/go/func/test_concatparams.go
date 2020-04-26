package main

import (
	"github.com/01-edu/z01"

	correct "./correct"
	student "./student"
)

func main() {
	table := [][]string{{"Hello", "how", "are", "you?"}}

	// 30 random slice of strings
	for i := 0; i < 30; i++ {
		table = append(table, z01.MultRandASCII())
	}
	for _, arg := range table {
		z01.Challenge("ConcatParams", student.ConcatParams, correct.ConcatParams, arg)
	}
}
