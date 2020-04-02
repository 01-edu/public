package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	if len(os.Args) <= 1 || len(os.Args) > 4 {
		z01.PrintRune('\n')
	} else if len(os.Args) == 4 {
		array := []rune(os.Args[1])
		search := []rune(os.Args[2])
		replace := []rune(os.Args[3])
		for _, val := range array {
			if val == search[0] {
				z01.PrintRune(replace[0])
			} else {
				z01.PrintRune(val)
			}
		}
		z01.PrintRune('\n')
	}
}
