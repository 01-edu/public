package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 2 {
		words := strings.Split(os.Args[1], " ")
		for i := len(words) - 1; i >= 0; i-- {
			if words[i] != "" {
				fmt.Println(words[i])
				return
			}
		}
	}
	fmt.Println()
}
