package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func IsNumeric(str string) bool {

	for i := 0; i < len(str); i++ {
		if !(str[i] >= '0' && str[i] <= '9') {
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
				z01.PrintRune(rune(number))
			} else if number <= 26 && number >= 1 && boole {
				number += 64
				z01.PrintRune(rune(number))
			} else {
				z01.PrintRune(' ')
			}
		} else {
			if !(arguments[i] == "--upper" && i == 0) {
				z01.PrintRune(' ')
			}
		}
	}
	z01.PrintRune('\n')
}
