package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(str string) {
	arrayStr := []rune(str)

	for i := 0; i < len(arrayStr); i++ {
		z01.PrintRune(arrayStr[i])
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {
	return nbr%2 == 0
}

func main() {
	EvenMsg := "I have an even number of arguments"
	OddMsg := "I have an odd number of arguments"

	lenOfArg := len(os.Args) - 1

	if isEven(lenOfArg) == true {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}
