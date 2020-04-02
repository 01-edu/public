package main

import (
	"fmt"
	"os"
	"strings"
)

func isAlphaNumerical(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9')
}

func isLowerRune(r rune) bool {
	return r >= 'a' && r <= 'z'
}

func toUpperRune(r rune) rune {
	if r >= 'a' && r <= 'z' {
		return r - 32
	}
	return r
}

func capitalize(s string) string {
	r := []rune(strings.ToLower(s))

	if isLowerRune(r[0]) {
		r[0] = toUpperRune(r[0])
	}

	for i := 1; i < len(r); i++ {
		if (!isAlphaNumerical(r[i-1])) && (isLowerRune(r[i])) {
			r[i] = toUpperRune(r[i])
		}
	}
	return string(r)
}

func main() {
	if len(os.Args) == 2 {
		fmt.Println(capitalize(os.Args[1]))
	} else if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
	} else {
		fmt.Println()
	}
}
