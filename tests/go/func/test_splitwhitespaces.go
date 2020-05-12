package main

import (
	"strings"

	"../lib"
	"./correct"
	"./student"
)

func main() {
	table := []string{}
	// 30 random slice of strings

	for i := 0; i < 30; i++ {
		val := strings.Join(lib.MultRandASCII(), " ")
		table = append(table, val)
	}

	table = append(table,
		"Hello how are you?")

	for _, arg := range table {
		lib.Challenge("SplitWhiteSpaces", student.SplitWhiteSpaces, correct.SplitWhiteSpaces, arg)
	}
}
