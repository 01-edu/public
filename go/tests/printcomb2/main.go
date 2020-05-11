package main

import (
	"fmt"

	"./student"
)

func printComb2() {
	a := 0
	b := 1
	for a <= 98 {
		b = a + 1
		for b <= 99 {
			fmt.Printf("%.2d %.2d", a, b)
			if !(a == 98 && b == 99) {
				fmt.Print(", ")
			}
			b++
		}
		a++
	}
	fmt.Println()
}

func main() {
	lib.Challenge("PrintComb2", student.PrintComb2, printComb2)
}
