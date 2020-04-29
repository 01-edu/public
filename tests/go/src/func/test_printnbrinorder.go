package main

import (
	"../lib"
	"./correct"
	"./student"
)

func main() {
	table := append(
		lib.MultRandIntBetween(0, lib.MaxInt),
		lib.MaxInt,
		321,
		321321,
		0,
	)
	for _, arg := range table {
		lib.Challenge("PrintNbrInOrder", student.PrintNbrInOrder, correct.PrintNbrInOrder, arg)
	}
}
