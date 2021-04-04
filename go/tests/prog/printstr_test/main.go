package main

import (
	"fmt"

	student "student"

	"github.com/01-edu/public/go/tests/lib"
)

func printStr(s string) {
	fmt.Print(s)
}

func main() {
	table := append(lib.MultRandASCII(), "Hello World!")
	for _, arg := range table {
		lib.Challenge("PrintStr", student.PrintStr, printStr, arg)
	}
}
