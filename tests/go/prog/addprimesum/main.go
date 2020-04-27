package main

import (
	"fmt"
	"os"
	"strconv"
)

func isPrime(nb int) bool {
	if nb <= 0 || nb == 1 {
		return false
	}
	if nb <= 3 {
		return true
	}
	if nb%2 == 0 || nb%3 == 0 {
		return false
	}

	i := 5
	for i*i <= nb {
		if nb%i == 0 || nb%(i+2) == 0 {
			return false
		}
		i += 6
	}
	return true
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("0")
	} else {
		argument, _ := strconv.Atoi(os.Args[1])

		if argument < 0 {
			fmt.Println("0")
		} else {
			result := 0
			for ; argument >= 0; argument-- {
				if isPrime(argument) {
					result += argument
				}
			}
			fmt.Println(result)
		}
	}
}
