package main

import (
	"math/bits"

	"../lib"
	"./correct"
	"./student"
)

func main() {
	table := append(
		lib.MultRandInt(),
		lib.IntRange(0, 12)...,
	)
	if bits.UintSize == 64 {
		table = append(table, lib.IntRange(13, 20)...)
	}
	for _, arg := range table {
		lib.Challenge("IterativeFactorial", student.IterativeFactorial, correct.IterativeFactorial, arg)
	}
}
