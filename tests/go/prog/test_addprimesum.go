package main

import (
	"strconv"

	"../lib"
)

func isAPrime(nb int) bool {
	if nb <= 0 || nb == 1 {
		return false
	}
	if nb <= 3 {
		return true
	}
	if nb%2 == 0 || nb%3 == 0 {
		return false
	}

	i := 5
	for i*i <= nb {
		if nb%i == 0 || nb%(i+2) == 0 {
			return false
		}
		i += 6
	}
	return true
}

func main() {
	// adds random numbers
	table := lib.MultRandIntBetween(1, 10000)

	// fill with all prime numbers between 0 and 100
	for i := 0; i < 100; i++ {
		if isAPrime(i) {
			table = append(table, i)
		}
	}

	for _, i := range table {
		lib.ChallengeMain("addprimesum", strconv.Itoa(i))
	}
	// special cases
	lib.ChallengeMain("addprimesum")
	lib.ChallengeMain("addprimesum", `""`)
	lib.ChallengeMain("addprimesum", "1", "2")
}
