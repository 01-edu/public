package main

import (
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	var table []string
	for i := 1; i < 10; i++ {
		table = append(table, strconv.Itoa(i))
	}
	for i := 0; i < 5; i++ {
		table = append(table, strconv.Itoa(z01.RandIntBetween(1, 1000)))
	}

	for _, arg := range table {
		z01.ChallengeMain(arg)
	}
}
