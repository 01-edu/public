package main

import (
	"strconv"

	"github.com/01-edu/z01"

	correct "./correct"
	student "./student"
)

func main() {
	table := make([]string, 30)
	for i := range table {
		table[i] = strconv.Itoa(z01.RandPosZ())
	}
	table = append(table,
		strconv.Itoa(z01.MaxInt),
		"",
		"0",
		"Invalid123",
		"123Invalid",
		"Invalid",
		"1Invalid23",
		"123",
		"123.",
		"123.0",
	)
	for _, arg := range table {
		z01.Challenge("BasicAtoi2", student.BasicAtoi2, correct.BasicAtoi2, arg)
	}
}
