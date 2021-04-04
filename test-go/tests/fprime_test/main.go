package main

import (
	"strconv"

	"github.com/01-edu/public/test-go/lib"
)

func main() {
	table := []string{
		" ",
		"1",
		"1 1",
		"hello",
		"p 1",
		"804577",
		"225225",
		"8333325",
		"42",
		"9539",
		"1000002",
		"1000003",
	}
	for i := 0; i < 10; i++ {
		table = append(table, strconv.Itoa(lib.RandIntBetween(1, 100)))
	}
	for _, s := range table {
		lib.ChallengeMain("fprime", s)
	}
}
