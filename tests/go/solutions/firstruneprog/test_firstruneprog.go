package main

import (
	"github.com/01-edu/z01"

	correct "./correct"
)

func main() {
	table := append(
		z01.MultRandASCII(),
		"Hello!",
		"Salut!",
		"Ola!",
		"♥01",
	)
	for _, arg := range table {
		z01.Challenge("FirstRuneProg", FirstRune, correct.FirstRune, arg)
	}
}
