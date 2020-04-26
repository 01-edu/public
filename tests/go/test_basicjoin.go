package main

import (
	"github.com/01-edu/z01"

	solutions "./solutions"
	student "./student"
)

func main() {
	table := [][]string{}

	// 30 valid pair of ramdom slice of strings to concatenate
	for i := 0; i < 30; i++ {
		table = append(table, z01.MultRandASCII())
	}
	table = append(table,
		[]string{"Hello!", " How are you?", "well and yourself?"},
	)
	for _, arg := range table {
		z01.Challenge("BasicJoin", student.BasicJoin, solutions.BasicJoin, arg)
	}
}
