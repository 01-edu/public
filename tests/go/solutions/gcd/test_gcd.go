package main

import (
	"strconv"

	"github.com/01-edu/z01"
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
		number1 := strconv.Itoa(z01.RandIntBetween(1, 100000))
		number2 := strconv.Itoa(z01.RandIntBetween(1, 100))
		args = append(args, []string{number1, number2})
	}
	for _, v := range args {
		z01.ChallengeMain("gcd", v...)
	}
}
