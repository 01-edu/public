package main

import (
	"github.com/01-edu/z01"
)

func main() {
	dif := 'a' - 'A'
	for c := 'a'; c <= 'z'; c++ {
		if c%2 == 0 {
			z01.PrintRune(c - dif)
		} else {
			z01.PrintRune(c)
		}
	}
	z01.PrintRune('\n')
}
