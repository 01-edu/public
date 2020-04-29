package main

import (
	"../lib"
	"./correct"
	"./student"
)

func main() {
	lib.Challenge("PrintComb2", student.PrintComb2, correct.PrintComb2)
}
