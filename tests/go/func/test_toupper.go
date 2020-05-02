package main

import (
	"../lib"
	"./correct"
	"./student"
)

func main() {
	table := lib.MultRandASCII()

	table = append(table,
		"Hello! How are you?",
	)

	for _, arg := range table {
		lib.Challenge("ToUpper", student.ToUpper, correct.ToUpper, arg)
	}
}
