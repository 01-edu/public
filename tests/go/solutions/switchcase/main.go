package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println()
	} else {
		array := []byte(os.Args[1])

		for i := 0; i < len(array); i++ {
			if array[i] >= 'A' && array[i] <= 'Z' {
				array[i] += 32
				z01.PrintRune(rune(array[i]))
			} else if array[i] >= 'a' && array[i] <= 'z' {
				array[i] -= 32
				z01.PrintRune(rune(array[i]))
			} else {
				z01.PrintRune(rune(array[i]))
			}
		}
		z01.PrintRune('\n')
	}
}
