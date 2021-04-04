package main

import (
	"strconv"

	student "student"

	"github.com/01-edu/public/test-go/lib"
)

func basicAtoi2(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func main() {
	table := make([]string, 30)
	for i := range table {
		table[i] = strconv.Itoa(lib.RandPosZ())
	}
	table = append(table,
		strconv.Itoa(lib.MaxInt),
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
		lib.Challenge("BasicAtoi2", student.BasicAtoi2, basicAtoi2, arg)
	}
}
