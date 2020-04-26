package main

import (
	"../lib"
	"./correct"
	"./student"
)

func main() {
	table := append(
		lib.MultRandInt(),
		lib.MinInt,
		lib.MaxInt,
		0,
	)
	for _, arg := range table {
		lib.Challenge("PrintNbr", student.PrintNbr, correct.PrintNbr, arg)
	}
}
