package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		z01.PrintRune('\n')
	} else {
		arrayRune := []byte(os.Args[1])

		for i := 0; i < len(arrayRune); i++ {
			if arrayRune[i] >= 'a' && arrayRune[i] <= 'z' {
				if arrayRune[i] >= ('a' + 13) {
					arrayRune[i] = arrayRune[i] - 13
				} else {
					arrayRune[i] = arrayRune[i] + 13
				}
			} else if arrayRune[i] >= 'A' && arrayRune[i] <= 'Z' {
				if arrayRune[i] >= ('A' + 13) {
					arrayRune[i] = arrayRune[i] - 13
				} else {
					arrayRune[i] = arrayRune[i] + 13
				}
			}
			z01.PrintRune(rune(arrayRune[i]))
		}
		z01.PrintRune('\n')
	}
}
