package main

import (
	"fmt"
	"os"
	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 || os.Args[1] == "" {
		z01.PrintRune('\n')
		return
	}
	letters := []rune(os.Args[1])
	switch letters[0] {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		fmt.Print(os.Args[1] + "ay\n")
	default:
		start := firstNotVowel(letters)
		if start == 0 {
			fmt.Println("No vowels")
		} else {
			fmt.Print(string(letters[start:]) + string(letters[:start]) + "ay\n")
		}
	}
}

func firstNotVowel(letters []rune) int {
	index := 0
	for i := 0; i < len(letters); i++ {
		if letters[i] == 'a' || letters[i] == 'e' || letters[i] == 'i' || letters[i] == 'o' || letters[i] == 'u' ||
			letters[i] == 'A' || letters[i] == 'E' || letters[i] == 'I' || letters[i] == 'O' || letters[i] == 'U' {
			return index
		}
		index++
	}
	return 0
}
