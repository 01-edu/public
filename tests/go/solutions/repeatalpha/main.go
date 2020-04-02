package main

import (
	"github.com/01-edu/z01"
	"os"
)

func toLowerCase(a rune) rune {
	if a >= 'A' && a <= 'Z' {
		dif := 'a' - 'A'
		return a + dif
	}
	return a
}

func main() {

	if len(os.Args) > 1 {
		arg := []rune(os.Args[1])
		for _, c := range arg {
			z01.PrintRune(c)
			rep := int(toLowerCase(c) - 'a')
			for i := 0; i < rep; i++ {
				z01.PrintRune(c)
			}
		}
	}

	z01.PrintRune('\n')
}
