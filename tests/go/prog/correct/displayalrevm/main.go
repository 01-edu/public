package main

import "fmt"

func main() {
	diff := 'a' - 'A'
	for c := 'z'; c >= 'a'; c-- {
		if c%2 == 0 {
			fmt.Printf("%c", c)
		} else {
			fmt.Printf("%c", c-diff)
		}
	}
	fmt.Println()
}
