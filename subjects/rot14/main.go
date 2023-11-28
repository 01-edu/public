package main

import (
	"piscine"

	"github.com/01-edu/z01"
)

func main() {
	result := piscine.Rot14("Hello! How are You?")

	for _, r := range result {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
