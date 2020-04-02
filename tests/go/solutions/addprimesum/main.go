package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func isPrime(nb int) bool {
	if nb <= 0 || nb == 1 {
		return false
	}
	if nb <= 3 {
		return true

	} else if nb%2 == 0 || nb%3 == 0 {
		return false
	}

	i := 5
	for i*i <= nb {
		if nb%i == 0 || nb%(i+2) == 0 {
			return false
		}
		i = i + 6
	}
	return true
}

func main() {
	if len(os.Args) != 2 {
		z01.PrintRune('0')
		z01.PrintRune('\n')
	} else {
		argument, _ := strconv.Atoi(os.Args[1])

		if argument < 0 {
			z01.PrintRune('0')
			z01.PrintRune('\n')
		} else {
			result := 0
			for ; argument >= 0; argument-- {
				if isPrime(argument) == true {
					result = result + argument
				}

			}
			fmt.Println(result)
		}
	}
}
