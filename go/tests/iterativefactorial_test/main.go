package main

import (
	"math/bits"

	student "student"

	"github.com/01-edu/public/go/tests/lib"
)

func iterativeFactorial(nb int) int {
	limit := 20
	if nb < 0 || nb > limit {
		return 0
	}
	if nb == 0 {
		return 1
	}
	return nb * iterativeFactorial(nb-1)
}

func main() {
	if bits.UintSize != 64 {
		panic("only works on 64 bits CPU")
	}
	table := append(
		lib.MultRandInt(),
		lib.IntRange(0, 20)...,
	)
	for _, arg := range table {
		lib.Challenge("IterativeFactorial", student.IterativeFactorial, iterativeFactorial, arg)
	}
}
