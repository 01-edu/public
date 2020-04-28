package main

import "fmt"

func main() {
	for r := 'z'; r >= 'a'; r-- {
		fmt.Print(r)
	}
	fmt.Println()
}
