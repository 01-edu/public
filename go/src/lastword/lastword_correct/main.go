package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 2 {
		words := strings.Fields(os.Args[1])
		if len(words) > 0 {
			fmt.Println(words[len(words)-1])
		}
	}
}
