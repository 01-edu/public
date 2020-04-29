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
		"12345",
		"0000012345",
		"000000",
	)
	for _, arg := range table {
		lib.Challenge("BasicAtoi", student.BasicAtoi, correct.BasicAtoi, arg)
	}
}
