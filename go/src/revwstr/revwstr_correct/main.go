package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 2 {
		a := strings.Split(os.Args[1], " ")
		for i := len(a) - 1; i >= 0; i-- {
			fmt.Print(a[i])
			if i != 0 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
