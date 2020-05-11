package main

import (
	"fmt"

	"./student"
)

func printStr(s string) {
	fmt.Print(s)
}

func main() {
	table := lib.MultRandASCII()
	table = append(table, "Hello World!")
	for _, arg := range table {
		lib.Challenge("PrintStr", student.PrintStr, printStr, arg)
	}
}
