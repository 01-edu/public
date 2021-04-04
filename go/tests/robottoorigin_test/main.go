package main

import (
	"github.com/01-edu/public/go/tests/lib"
)

func main() {
	table := []string{
		"UD",
		"LL",
	}

	for i := 0; i < 15; i++ {
		table = append(table, lib.RandStr(lib.RandIntBetween(5, 1000), "UDLR"))
	}

	for _, arg := range table {
		lib.ChallengeMain("robottoorigin", arg)
	}
}
