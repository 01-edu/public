package main

import (
	"fmt"

	"./student"
	"github.com/01-edu/public/go/lib"
)

func main() {
	table := append(lib.MultRandASCII(), "Hello World!")
	for _, arg := range table {
		lib.Challenge("PrintStr", student.PrintStr, fmt.Print, arg)
	}
}
