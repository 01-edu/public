package main

import (
	"github.com/01-edu/z01"

	correct "./correct"
	student "./student"
)

func main() {
	a := append(z01.MultRandIntBetween(0, 99999), 5, 4, 1)
	for _, elem := range a {
		z01.Challenge("FindPrevPrime", student.FindPrevPrime, correct.FindPrevPrime, elem)
	}
}
