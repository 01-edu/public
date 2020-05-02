package main

import (
	"../lib"
	"./correct"
	"./student"
)

func main() {
	table := []string{
		" ",
		"Hello 78 World!    4455 /",
	}
	for l := 0; l < 7; l++ {
		a := lib.RandIntBetween(5, 20)
		b := lib.RandASCII()
		table = append(table, lib.RandStr(a, b))
	}

	for _, arg := range table {
		lib.Challenge("AlphaCount", student.AlphaCount, correct.AlphaCount, arg)
	}
}
