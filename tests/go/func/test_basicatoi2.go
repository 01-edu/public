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
		lib.Challenge("BasicAtoi2", student.BasicAtoi2, correct.BasicAtoi2, arg)
	}
}
