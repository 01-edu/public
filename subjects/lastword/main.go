package main

import (
	"piscine"

	"github.com/01-edu/z01"
)

func main() {
	z01.PrintRune(piscine.LastWord("this        ...       is sparta, then again, maybe    not"))
	z01.PrintRune(piscine.LastWord(" "))
	z01.PrintRune(piscine.LastWord(" lorem,ipsum "))
}
