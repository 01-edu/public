package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 4 {
		fmt.Println(strings.Map(func(r rune) rune {
			if r == []rune(os.Args[2])[0] {
				return []rune(os.Args[3])[0]
			}
			return r
		}, os.Args[1]))
	}
}
