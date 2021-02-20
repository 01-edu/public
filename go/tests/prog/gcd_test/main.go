package main

import (
	"strconv"

	"github.com/01-edu/public/go/tests/lib"
)

func main() {
	args := [][]string{
		{"23"},
		{"12", "23"},
		{"25", "15"},
		{"23043", "122"},
		{"11", "77"},
	}
	for i := 0; i < 25; i++ {
		a := strconv.Itoa(lib.RandIntBetween(1, 100000))
		b := strconv.Itoa(lib.RandIntBetween(1, 100))
		args = append(args, []string{a, b})
	}
	for _, v := range args {
		lib.ChallengeMain("gcd", v...)
	}
}
