package main

import (
	"strings"

	student "student"

	"github.com/01-edu/public/test-go/lib"
)

func concatParams(args []string) string {
	return strings.Join(args, "\n")
}

func main() {
	table := [][]string{{"Hello", "how", "are", "you?"}}

	// 30 random slice of strings
	for i := 0; i < 30; i++ {
		table = append(table, lib.MultRandASCII())
	}
	for _, arg := range table {
		lib.Challenge("ConcatParams", student.ConcatParams, concatParams, arg)
	}
}
