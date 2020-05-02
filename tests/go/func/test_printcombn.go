package main

import (
	"../lib"
	"./correct"
	"./student"
)

func main() {
	table := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, arg := range table {
		lib.Challenge("PrintCombN", student.PrintCombN, correct.PrintCombN, arg)
	}
}
