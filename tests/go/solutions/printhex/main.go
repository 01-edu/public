package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func printBase(nbr int, base string) {
	var result []rune
	typeBase := []rune(base)
	baseSize := len(base)
	pos := 0
	for nbr > 0 {
		pos = nbr % baseSize
		result = append(result, typeBase[pos])
		nbr = nbr / baseSize
	}
	for j := len(result) - 1; j >= 0; j-- {
		z01.PrintRune(result[j])
	}
}

func main() {
	if len(os.Args) != 2 {
		z01.PrintRune('\n')
	} else {
		nbr, _ := strconv.Atoi(os.Args[1])
		if nbr != 0 {
			printBase(nbr, "0123456789abcdef")
			z01.PrintRune('\n')
		} else {
			fmt.Println(0)
		}
	}
}
