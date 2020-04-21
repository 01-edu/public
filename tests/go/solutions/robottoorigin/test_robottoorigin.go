package main

import (
	"github.com/01-edu/z01"
)

func main() {
	table := []string{}

	table = append(table,
		"UD",
		"LL",
	)

	for i := 0; i < 15; i++ {
		value := z01.RandStr(z01.RandIntBetween(5, 1000), "UDLR")
		table = append(table, value)
	}

	for _, arg := range table {
		z01.ChallengeMain("robottoorigin", arg)
	}
}
