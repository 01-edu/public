package main

import (
	"strconv"
	"strings"

	"lib"
)

func main() {
	table := []string{
		" ",
		"123 132 1",
		"1 5",
		"0",
	}
	for i := 0; i < 10; i++ {
		table = append(table, strconv.Itoa(lib.RandIntBetween(-1000, lib.MaxInt)))
	}
	for i := 0; i < 15; i++ {
		table = append(table, strconv.Itoa(i))
	}
	for _, s := range table {
		lib.ChallengeMain("printhex", strings.Fields(s)...)
	}
}
