package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 2 {
		for i := 0; i < len(os.Args[1]); i++ {
			if os.Args[1][i] != ' ' {
				z01.PrintRune(rune(os.Args[1][i]))

				if i != len(os.Args[1])-1 {
					if os.Args[1][i+1] == ' ' {
						z01.PrintRune('\n')
						return
					}
				}
			}
		}
		z01.PrintRune('\n')
	} else {
		fmt.Println()
	}
}
