package main

import (
	"github.com/01-edu/z01"
	"strconv"
	"testing"
)

func TestPrintChessBoard(t *testing.T) {
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
		number1 := strconv.Itoa(z01.RandIntBetween(1, 200))
		number2 := strconv.Itoa(z01.RandIntBetween(1, 200))
		table = append(table, []string{number1, number2})
	}

	for _, v := range table {
		z01.ChallengeMainExam(t, v...)
	}
}
