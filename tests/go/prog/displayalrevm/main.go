package main

import "github.com/01-edu/z01"

func main() {
	diff := 'a' - 'A'
	for c := 'z'; c >= 'a'; c-- {
		if c%2 == 0 {
			z01.PrintRune(c)
		} else {
			z01.PrintRune(c - diff)
		}
	}
	z01.PrintRune('\n')
}
