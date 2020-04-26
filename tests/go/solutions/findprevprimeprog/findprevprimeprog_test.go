package main

import (
	"github.com/01-edu/z01"

	solutions "../../solutions"
)

func main() {
	a := append(z01.MultRandIntBetween(0, 99999), 5, 4, 1)
	for _, elem := range a {
		z01.Challenge("FindPrevPrimeProg", FindPrevPrime, solutions.FindPrevPrime, elem)
	}
}
