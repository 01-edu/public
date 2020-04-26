package main

import (
	"os"
	"unicode"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 2 {
		for _, r := range os.Args[1] {
			z01.PrintRune(r)
			if unicode.IsLetter(r) {
				rep := unicode.ToLower(r) - 'a'
				for i := 0; i < int(rep); i++ {
					z01.PrintRune(r)
				}
			}
		}
		z01.PrintRune('\n')
	}
}
