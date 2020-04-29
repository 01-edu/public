package main

import (
	"../lib"
	"./correct"
	"./student"
)

func main() {
	table := append(
		lib.MultRandASCII(),
		"Hello!",
		"Salut!",
		"Ola!",
		"♥01",
	)
	for _, arg := range table {
		lib.Challenge("FirstRune", student.FirstRune, correct.FirstRune, arg)
	}
}
