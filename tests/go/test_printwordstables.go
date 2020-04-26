package main

import (
	"strings"

	"github.com/01-edu/z01"

	correct "./correct"
	student "./student"
)

func main() {
	table := [][]string{}
	// 30 random slice of slice of strings

	for i := 0; i < 30; i++ {
		val := correct.SplitWhiteSpaces(strings.Join(z01.MultRandASCII(), " "))
		table = append(table, val)
	}

	table = append(table,
		[]string{"Hello", "how", "are", "you?"})

	for _, arg := range table {
		z01.Challenge("PrintWordsTables", student.PrintWordsTables, correct.PrintWordsTables, arg)
	}
}
