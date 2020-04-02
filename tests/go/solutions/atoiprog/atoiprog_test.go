package main

import (
	"strconv"
	"testing"

	"github.com/01-edu/z01"

	solutions "../../solutions"
)

func TestAtoiProg(t *testing.T) {
	table := make([]string, 30)
	for i := range table {
		table[i] = strconv.Itoa(z01.RandInt())
	}
	table = append(table,
		strconv.Itoa(z01.MinInt),
		strconv.Itoa(z01.MaxInt),
		"",
		"-",
		"+",
		"0",
		"+0",
		"-Invalid123",
		"--123",
		"-+123",
		"++123",
		"123-",
		"123+",
		"123.",
		"123.0",
		"123a45",
	)
	for _, arg := range table {
		z01.Challenge(t, Atoi, solutions.Atoi, arg)
	}
}
