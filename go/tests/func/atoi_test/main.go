package main

import (
	"strconv"

	student "student"

	"github.com/01-edu/public/go/tests/lib"
)

func atoi(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func main() {
	table := make([]string, 30)
	for i := range table {
		table[i] = strconv.Itoa(lib.RandInt())
	}
	table = append(table,
		strconv.Itoa(lib.MinInt),
		strconv.Itoa(lib.MaxInt),
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
		lib.Challenge("Atoi", student.Atoi, atoi, arg)
	}
}
