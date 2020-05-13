package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 2 {
		for _, r := range os.Args[1] {
			if r >= 'a' && r <= 'z' {
				if r >= 'a'+13 {
					r -= 13
				} else {
					r += 13
				}
			} else if r >= 'A' && r <= 'Z' {
				if r >= 'A'+13 {
					r -= 13
				} else {
					r += 13
				}
			}
			fmt.Printf("%c", r)
		}
		fmt.Println()
	}
}
