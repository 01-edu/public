package main

import (
	"../lib"
	"./correct"
	"./student"
)

func main() {
	table := lib.MultRandASCII()
	table = append(table, "Hello World!")
	for _, arg := range table {
		lib.Challenge("PrintStr", student.PrintStr, correct.PrintStr, arg)
	}
}
