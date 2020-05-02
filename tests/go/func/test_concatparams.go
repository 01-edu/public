package main

import (
	"../lib"
	"./correct"
	"./student"
)

func main() {
	table := [][]string{{"Hello", "how", "are", "you?"}}

	// 30 random slice of strings
	for i := 0; i < 30; i++ {
		table = append(table, lib.MultRandASCII())
	}
	for _, arg := range table {
		lib.Challenge("ConcatParams", student.ConcatParams, correct.ConcatParams, arg)
	}
}
