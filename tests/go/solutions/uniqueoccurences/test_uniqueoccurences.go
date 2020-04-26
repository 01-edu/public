package main

import "github.com/01-edu/z01"

func main() {
	table := []string{
		"abbaac",
		"ab",
		"abcacccazb",
		"",
	}

	for i := 0; i < 15; i++ {
		table = append(table, z01.RandStr(z01.RandIntBetween(5, 10), z01.Lower))
	}

	for _, arg := range table {
		z01.ChallengeMain("uniqueoccurences", arg)
	}
}
