package main

import (
	"../lib"
	"./correct"
	"./student"
)

func main() {
	a := append(lib.MultRandIntBetween(0, 99999), 5, 4, 1)
	for _, elem := range a {
		lib.Challenge("FindPrevPrime", student.FindPrevPrime, correct.FindPrevPrime, elem)
	}
}
