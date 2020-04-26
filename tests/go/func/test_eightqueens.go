package main

import (
	"../lib"
	"./correct"
	"./student"
)

func main() {
	lib.Challenge("EightQueens", student.EightQueens, correct.EightQueens)
}
