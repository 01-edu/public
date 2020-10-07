package main

import "github.com/01-edu/z01"

func main() {
	alphabet := "abcdefghijklmnopqrstuvwxyz\n"
	for _, letter := range alphabet {
		z01.PrintRune(letter)
	}
}
