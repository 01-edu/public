package main

import (
	"math/bits"

	"github.com/01-edu/z01"

	solutions "./solutions"
	student "./student"
)

func main() {
	table := append(
		z01.MultRandInt(),
		z01.IntRange(0, 12)...,
	)
	if bits.UintSize == 64 {
		table = append(table, z01.IntRange(13, 20)...)
	}
	for _, arg := range table {
		z01.Challenge("RecursiveFactorial", student.RecursiveFactorial, solutions.RecursiveFactorial, arg)
	}
}
