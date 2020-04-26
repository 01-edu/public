package main

import (
	"fmt"
	"os"
	"unicode"
)

func main() {
	if len(os.Args) == 2 {
		s := os.Args[1]
		for i, r := range s {
			if unicode.IsLower(r) {
				s[i] = unicode.ToUpper(r)
			} else if unicode.IsUpper(r) {
				s[i] = unicode.ToLower(r)
			}
		}
		fmt.Println(s)
	}
}
