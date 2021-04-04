package main

import (
	"strconv"

	student "student"

	"github.com/01-edu/public/test-go/lib"
)

func basicAtoi(s string) int {
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
		"12345",
		"0000012345",
		"000000",
	)
	for _, arg := range table {
		lib.Challenge("BasicAtoi", student.BasicAtoi, basicAtoi, arg)
	}
}
