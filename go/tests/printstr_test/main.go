package main

import (
	"fmt"

	student "student"

	"lib"
)

func main() {
	table := append(lib.MultRandASCII(), "Hello World!")
	for _, arg := range table {
		lib.Challenge("PrintStr", student.PrintStr, fmt.Print, arg)
	}
}
