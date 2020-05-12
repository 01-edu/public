package main

import (
	"fmt"
	"os"
	"unicode"
)

func main() {
	for _, arg := range os.Args[1:] {
		arg := []rune(arg)
		for j, r := range arg {
			if j+1 == len(arg) || arg[j+1] == ' ' {
				arg[j] = unicode.ToUpper(r)
			} else {
				arg[j] = unicode.ToLower(r)
			}
		}
		fmt.Println(string(arg))
	}
}
