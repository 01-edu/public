package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Print(piscine.FirstWord("hello there"))
	fmt.Print(piscine.FirstWord(""))
	fmt.Print(piscine.FirstWord("hello   .........  bye"))
}
