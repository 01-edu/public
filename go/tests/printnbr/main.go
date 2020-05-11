package main

import (
	"fmt"

	"./student"
	"github.com/01-edu/public/go/lib"
)

func printNbr(n int) {
	fmt.Print(n)
}

func main() {
	table := append(
		lib.MultRandInt(),
		lib.MinInt,
		lib.MaxInt,
		0,
	)
	for _, arg := range table {
		lib.Challenge("PrintNbr", student.PrintNbr, printNbr, arg)
	}
}
