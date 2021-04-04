package main

import (
	student "student"

	"github.com/01-edu/public/go/tests/lib"
)

func lastRune(s string) rune {
	runes := []rune(s)
	index := len(runes) - 1
	return runes[index]
}

func main() {
	table := lib.MultRandASCII()
	table = append(table,
		"Hello!",
		"Salut!",
		"Ola!",
		lib.RandStr(lib.RandIntBetween(1, 15), lib.RandAlnum()),
	)
	for _, arg := range table {
		lib.Challenge("LastRune", student.LastRune, lastRune, arg)
	}
}
