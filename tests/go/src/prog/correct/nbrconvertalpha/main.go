package main

import (
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func IsNumeric(s string) bool {
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

func main() {
	arguments := os.Args[1:]
	for i := range arguments {
		if IsNumeric(arguments[i]) {
			number, _ := strconv.Atoi(arguments[i])
			boole := false
			if os.Args[1] == "--upper" {
				boole = true
			}

			if number <= 26 && number >= 1 && !boole {
				number += 96
				fmt.Printf("%c", rune(number))
			} else if number <= 26 && number >= 1 && boole {
				number += 64
				fmt.Printf("%c", rune(number))
			} else {
				fmt.Print(" ")
			}
		} else {
			if !(arguments[i] == "--upper" && i == 0) {
				fmt.Print(" ")
			}
		}
	}
	fmt.Println()
}
