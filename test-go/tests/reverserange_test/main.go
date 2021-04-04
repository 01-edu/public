package main

import (
	"strconv"

	"github.com/01-edu/public/test-go/lib"
)

func main() {
	for i := 0; i < 10; i++ {
		start := lib.RandIntBetween(-20, 20)
		end := lib.RandIntBetween(-20, 20)
		lib.ChallengeMain("reverserange", strconv.Itoa(start), strconv.Itoa(end))
	}
	lib.ChallengeMain("reverserange", "2", "1", "3")
	lib.ChallengeMain("reverserange", "a", "1")
	lib.ChallengeMain("reverserange", "1", "b")
	lib.ChallengeMain("reverserange", "1", "nan")
	lib.ChallengeMain("reverserange", "nan", "b")
	lib.ChallengeMain("reverserange")
}
