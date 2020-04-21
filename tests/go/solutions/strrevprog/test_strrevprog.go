package main

import (
	"github.com/01-edu/z01"
)

func main() {
	table := append(
		z01.MultRandASCII(),
		"Hello!",
		"Bonjour!",
		"Hola!",
	)
	for _, arg := range table {
		z01.ChallengeMain("strrevprog", arg)
	}
}
