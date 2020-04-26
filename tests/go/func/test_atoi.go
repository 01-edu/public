package main

import (
	"strconv"

	"../lib"
	"./correct"
	"./student"
)

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
		lib.Challenge("Atoi", student.Atoi, correct.Atoi, arg)
	}
}
