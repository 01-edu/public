package main

import (
	"fmt"

	"./student"
	"github.com/01-edu/public/go/lib"
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
