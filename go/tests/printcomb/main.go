package main

import (
	"fmt"

	"./student"
)

func printComb() {
	for i := '0'; i <= '7'; i++ {
		for j := i + 1; j <= '8'; j++ {
			for k := j + 1; k <= '9'; k++ {
				fmt.Printf("%c%c%c", i, j, k)
				if i < '7' {
					fmt.Printf(", ")
				} else {
					fmt.Println()
				}
			}
		}
	}
}

func main() {
	lib.Challenge("PrintComb", student.PrintComb, printComb)
}
