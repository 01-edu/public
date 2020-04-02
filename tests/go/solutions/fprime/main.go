package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func fprime(value int) {
	if value != 1 {
		divisionIterator := 2
		for value > 1 {
			if value%divisionIterator == 0 {
				fmt.Print(divisionIterator)
				value = value / divisionIterator

				if value > 1 {
					z01.PrintRune('*')
				}
				divisionIterator--
			}
			divisionIterator++
		}
	}

	z01.PrintRune('\n')
}

func main() {
	if len(os.Args) != 2 {
		z01.PrintRune('\n')
	} else {
		par, _ := strconv.Atoi(os.Args[1])
		fprime(par)
	}

}
