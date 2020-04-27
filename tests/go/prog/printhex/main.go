package main

import (
	"fmt"
	"os"
	"strconv"
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
		fmt.Printf("%c", result[j])
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println()
	} else {
		nbr, _ := strconv.Atoi(os.Args[1])
		if nbr != 0 {
			printBase(nbr, "0123456789abcdef")
			fmt.Println()
		} else {
			fmt.Println(0)
		}
	}
}
