package main

import (
	"strconv"

	"github.com/01-edu/z01"

	solutions "./solutions"
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
		"12345",
		"0000012345",
		"000000",
	)
	for _, arg := range table {
		z01.Challenge("BasicAtoi", student.BasicAtoi, solutions.BasicAtoi, arg)
	}
}
