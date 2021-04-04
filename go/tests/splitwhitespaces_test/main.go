package main

import (
	"strings"

	student "student"

	"github.com/01-edu/public/go/tests/lib"
)

func main() {
	table := []string{"Hello how are you?"}

	// 30 random slice of strings
	for i := 0; i < 30; i++ {
		table = append(table, strings.Join(lib.MultRandASCII(), " "))
	}
	for _, arg := range table {
		lib.Challenge("SplitWhiteSpaces", student.SplitWhiteSpaces, strings.Fields, arg)
	}
}
