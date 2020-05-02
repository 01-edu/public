package main

import (
	"fmt"
	"os"
	"unicode"
)

func main() {
	if len(os.Args) == 2 {
		s := os.Args[1]
		runes := []rune(s)
		for i, r := range s {
			if unicode.IsLower(r) {
				runes[i] = unicode.ToUpper(r)
			} else if unicode.IsUpper(r) {
				runes[i] = unicode.ToLower(r)
			}
		}
		fmt.Println(string(runes))
	}
}
