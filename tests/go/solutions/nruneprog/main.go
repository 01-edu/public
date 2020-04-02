package main

import (
	"fmt"
	"os"
	"strconv"

	solutions ".."
)

func NRune(s string, n int) rune {
	if n > len(s) || n < 1 {
		return 0
	}
	runes := []rune(s)
	return runes[n-1]
}

func main() {
	if len(os.Args) != 3 {
		return
	}
	val, err := strconv.Atoi(os.Args[2])

	if err != nil {
		fmt.Printf("\"%s\" is not an integer value\n", os.Args[2])
		return
	}

	rune := solutions.NRune(os.Args[1], val)

	if rune == 0 {
		fmt.Printf("Invalid position: \"%d\" in \"%s\"\n", val, os.Args[1])
		return
	}

	fmt.Printf("%c\n", rune)
}
