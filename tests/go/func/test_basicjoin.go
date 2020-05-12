package main

import (
	"../lib"
	"./correct"
	"./student"
)

func main() {
	table := [][]string{}

	// 30 valid pair of ramdom slice of strings to concatenate
	for i := 0; i < 30; i++ {
		table = append(table, lib.MultRandASCII())
	}
	table = append(table,
		[]string{"Hello!", " How are you?", "well and yourself?"},
	)
	for _, arg := range table {
		lib.Challenge("BasicJoin", student.BasicJoin, correct.BasicJoin, arg)
	}
}
