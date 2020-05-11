package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println()
	} else {
		arrString := strings.Split(os.Args[1], " ")
		for i := len(arrString) - 1; i >= 0; i-- {
			fmt.Print(arrString[i])
			if i != 0 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
