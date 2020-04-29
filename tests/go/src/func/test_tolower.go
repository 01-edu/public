package main

import (
	"../lib"
	"./correct"
	"./student"
)

func main() {
	table := append(
		lib.MultRandASCII(),
		"Hello! How are you?",
	)
	for _, arg := range table {
		lib.Challenge("ToLower", student.ToLower, correct.ToLower, arg)
	}
}
