package main

import "github.com/01-edu/z01"

func main() {
	table := []string{
		"UD",
		"LL",
	}

	for i := 0; i < 15; i++ {
		table = append(table, z01.RandStr(z01.RandIntBetween(5, 1000), "UDLR"))
	}

	for _, arg := range table {
		z01.ChallengeMain("robottoorigin", arg)
	}
}
