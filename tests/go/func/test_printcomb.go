package main

import (
	"../lib"
	"./correct"
	"./student"
)

func main() {
	lib.Challenge("PrintComb", student.PrintComb, correct.PrintComb)
}
