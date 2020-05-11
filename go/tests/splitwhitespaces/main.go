package main

import (
	"strings"

	"./student"
	"github.com/01-edu/public/go/lib"
)

func splitWhiteSpaces(s string) []string {
	return strings.Fields(s)
}

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
		lib.Challenge("SplitWhiteSpaces", student.SplitWhiteSpaces, splitWhiteSpaces, arg)
	}
}
