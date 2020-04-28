package main

import (
	"strconv"

	"../../lib"
)

func main() {
	var table []string
	for i := 1; i < 10; i++ {
		table = append(table, strconv.Itoa(i))
	}
	for i := 0; i < 5; i++ {
		table = append(table, strconv.Itoa(lib.RandIntBetween(1, 1000)))
	}

	for _, arg := range table {
		lib.ChallengeMain("tabmult", arg)
	}
}
