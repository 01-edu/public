package main

import (
	"fmt"
	"os"
	"unicode"
)

func main() {
	if len(os.Args) == 2 {
		runes := []rune(os.Args[1])
		for i, r := range runes {
			if unicode.IsLower(r) {
				runes[i] = unicode.ToUpper(r)
			} else if unicode.IsUpper(r) {
				runes[i] = unicode.ToLower(r)
			}
		}
		fmt.Println(string(runes))
	}
}
