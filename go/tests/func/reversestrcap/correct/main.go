package main

import (
	"fmt"
	"os"
)

func upperCase(ch rune) rune {
	if ch >= 'a' && ch <= 'z' {
		return ch - ('a' - 'A')
	}
	return ch
}

func lowerCase(ch rune) rune {
	if ch >= 'A' && ch <= 'Z' {
		return ch + ('a' - 'A')
	}
	return ch
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println()
		os.Exit(0)
	}

	for i := 1; i < len(os.Args); i++ {
		arg := []rune(os.Args[i])
		for j, c := range arg {
			if j+1 < len(arg) && arg[j+1] == ' ' {
				arg[j] = upperCase(c)
			} else if j+1 == len(arg) {
				arg[j] = upperCase(c)
			} else {
				arg[j] = lowerCase(c)
			}
		}
		fmt.Println(string(arg))
	}
}
