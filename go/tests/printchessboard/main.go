package main

import (
	"strconv"

	"github.com/01-edu/public/go/lib"
)

func main() {
	table := [][]string{}

	table = append(table,
		[]string{"0", "0"},
		[]string{"1", "2"},
		[]string{"2"},
		[]string{"0"},
		[]string{"3", "3"},
		[]string{"1", "5"},
		[]string{"5", "1"},
		[]string{"4", "3"},
	)

	for i := 0; i < 2; i++ {
		number1 := strconv.Itoa(lib.RandIntBetween(1, 200))
		number2 := strconv.Itoa(lib.RandIntBetween(1, 200))
		table = append(table, []string{number1, number2})
	}

	for _, v := range table {
		lib.ChallengeMain("printchessboard", v...)
	}
}
