package main

import (
	"../lib"
	"./correct"
	"./student"
)

func main() {
	table := append(
		lib.MultRandASCII(),
		"Hello! How are you? How+are+things+4you?",
		"Hello! How are you?",
		"a",
		"z",
		"!",
		"9a",
		"9a LALALA!",
	)
	for _, arg := range table {
		lib.Challenge("Capitalize", student.Capitalize, correct.Capitalize, arg)
	}
}
