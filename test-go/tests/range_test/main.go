package main

import (
	"strconv"

	"github.com/01-edu/public/test-go/lib"
)

func main() {
	for i := 0; i < 10; i++ {
		start := lib.RandIntBetween(-20, 20)
		end := lib.RandIntBetween(-20, 20)
		lib.ChallengeMain("range", strconv.Itoa(start), strconv.Itoa(end))
	}
	lib.ChallengeMain("range", "2", "1", "3")
	lib.ChallengeMain("range", "a", "1")
	lib.ChallengeMain("range", "1", "b")
	lib.ChallengeMain("range", "1", "nan")
	lib.ChallengeMain("range", "nan", "b")
	lib.ChallengeMain("range")
}
