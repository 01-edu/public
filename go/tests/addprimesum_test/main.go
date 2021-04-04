package main

import (
	"strconv"

	"github.com/01-edu/public/go/tests/lib"
	"github.com/01-edu/public/go/tests/lib/is"
)

func main() {
	// adds random numbers
	table := lib.MultRandIntBetween(1, 10000)

	// fill with all prime numbers between 0 and 100
	for i := 0; i < 100; i++ {
		if is.Prime(i) {
			table = append(table, i)
		}
	}

	for _, i := range table {
		lib.ChallengeMain("addprimesum", strconv.Itoa(i))
	}
	// special cases
	lib.ChallengeMain("addprimesum")
	lib.ChallengeMain("addprimesum", `""`)
	lib.ChallengeMain("addprimesum", "1", "2")
}
