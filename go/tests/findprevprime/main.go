package main

import (
	"./student"
	"github.com/01-edu/public/go/lib"
	"github.com/01-edu/public/go/lib/is"
)

func findPrevPrime(nb int) int {
	if nb < 2 {
		return 0
	}
	if is.Prime(nb) {
		return nb
	}
	return findPrevPrime(nb - 1)
}

func main() {
	a := append(lib.MultRandIntBetween(0, 99999), 5, 4, 1)
	for _, elem := range a {
		lib.Challenge("FindPrevPrime", student.FindPrevPrime, findPrevPrime, elem)
	}
}
