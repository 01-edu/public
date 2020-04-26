package main

import (
	"../lib"
	"./correct"
	"./student"
)

func main() {
	table := lib.MultRandASCII()
	table = append(table,
		"Hello!",
		"Salut!",
		"Ola!",
		lib.RandStr(lib.RandIntBetween(1, 15), lib.RandAlnum()),
	)
	for _, arg := range table {
		lib.Challenge("LastRune", student.LastRune, correct.LastRune, arg)
	}
}
