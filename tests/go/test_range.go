package main

import (
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	for i := 0; i < 10; i++ {
		start := z01.RandIntBetween(-20, 20)
		end := z01.RandIntBetween(-20, 20)
		z01.ChallengeMain("range", strconv.Itoa(start), strconv.Itoa(end))
	}
	z01.ChallengeMain("range", "2", "1", "3")
	z01.ChallengeMain("range", "a", "1")
	z01.ChallengeMain("range", "1", "b")
	z01.ChallengeMain("range", "1", "nan")
	z01.ChallengeMain("range", "nan", "b")
	z01.ChallengeMain("range")
}
