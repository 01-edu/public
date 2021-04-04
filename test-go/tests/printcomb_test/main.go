package main

import (
	"fmt"

	student "student"

	"github.com/01-edu/public/test-go/lib"
)

func printComb() {
	for i := '0'; i <= '7'; i++ {
		for j := i + 1; j <= '8'; j++ {
			for k := j + 1; k <= '9'; k++ {
				fmt.Printf("%c%c%c", i, j, k)
				if i < '7' {
					fmt.Print(", ")
				}
			}
		}
	}
	fmt.Println()
}

func main() {
	lib.Challenge("PrintComb", student.PrintComb, printComb)
}
