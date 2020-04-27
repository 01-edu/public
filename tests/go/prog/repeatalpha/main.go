package main

import (
	"fmt"
	"os"
	"unicode"
)

func main() {
	if len(os.Args) == 2 {
		for _, r := range os.Args[1] {
			fmt.Printf("%c", r)
			if unicode.IsLetter(r) {
				rep := unicode.ToLower(r) - 'a'
				for i := 0; i < int(rep); i++ {
					fmt.Printf("%c", r)
				}
			}
		}
		fmt.Println()
	}
}
