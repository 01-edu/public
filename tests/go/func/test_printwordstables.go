package main

import (
	"strings"

	"../lib"
	"./correct"
	"./student"
)

func main() {
	table := [][]string{}
	// 30 random slice of slice of strings

	for i := 0; i < 30; i++ {
		val := correct.SplitWhiteSpaces(strings.Join(lib.MultRandASCII(), " "))
		table = append(table, val)
	}

	table = append(table,
		[]string{"Hello", "how", "are", "you?"})

	for _, arg := range table {
		lib.Challenge("PrintWordsTables", student.PrintWordsTables, correct.PrintWordsTables, arg)
	}
}
