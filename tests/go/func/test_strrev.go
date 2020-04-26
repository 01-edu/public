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
		"Bonjour!",
		"Hola!",
	)
	for _, arg := range table {
		lib.Challenge("StrRev", student.StrRev, correct.StrRev, arg)
	}
}
