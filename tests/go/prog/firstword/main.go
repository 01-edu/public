package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 2 {
		for i := 0; i < len(os.Args[1]); i++ {
			if os.Args[1][i] != ' ' {
				fmt.Printf("%c", rune(os.Args[1][i]))

				if i != len(os.Args[1])-1 {
					if os.Args[1][i+1] == ' ' {
						fmt.Println()
						return
					}
				}
			}
		}
		fmt.Println()
	}
}
