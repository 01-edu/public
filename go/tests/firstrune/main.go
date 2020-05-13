package main

import (
	student "student"

	"lib"
)

func firstRune(s string) rune {
	runes := []rune(s)
	return runes[0]
}

func main() {
	table := append(
		lib.MultRandASCII(),
		"Hello!",
		"Salut!",
		"Ola!",
		"♥01",
	)
	for _, arg := range table {
		lib.Challenge("FirstRune", student.FirstRune, firstRune, arg)
	}
}
