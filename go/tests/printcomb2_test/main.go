package main

import (
	"fmt"

	student "student"

	"github.com/01-edu/public/go/tests/lib"
)

func printComb2() {
	for a := 0; a <= 98; a++ {
		for b := a + 1; b <= 99; b++ {
			fmt.Printf("%.2d %.2d", a, b)
			if a != 98 || b != 99 {
				fmt.Print(", ")
			}
		}
	}
	fmt.Println()
}

func main() {
	lib.Challenge("PrintComb2", student.PrintComb2, printComb2)
}
