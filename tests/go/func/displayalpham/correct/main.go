package main

import "fmt"

func main() {
	diff := 'a' - 'A'
	for c := 'a'; c <= 'z'; c++ {
		if c%2 == 0 {
			fmt.Printf("%c", c-diff)
		} else {
			fmt.Printf("%c", c)
		}
	}
	fmt.Println()
}
