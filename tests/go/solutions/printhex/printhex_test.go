package main

import (
	"strconv"
	"strings"
	"testing"

	"github.com/01-edu/z01"
)

func TestPrintHex(t *testing.T) {
	var table []string
	table = append(table,
		" ",
		"123 132 1",
		"1 5",
		"0",
	)
	for i := 0; i < 10; i++ {
		table = append(table, strconv.Itoa(z01.RandIntBetween(-1000, 2000000000)))
	}

	for i := 0; i < 15; i++ {
		table = append(table, strconv.Itoa(i))
	}

	for _, s := range table {
		z01.ChallengeMainExam(t, strings.Fields(s)...)
	}
}
