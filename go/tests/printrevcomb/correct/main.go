package main

import "fmt"

func main() {
	for a := 9; a >= 2; a-- {
		for b := a - 1; b >= 1; b-- {
			for c := b - 1; c >= 0; c-- {
				fmt.Printf("%d%d%d", a, b, c)
				if a+b+c != 3 {
					fmt.Print(", ")
				}
			}
		}
	}
	fmt.Println()
}
