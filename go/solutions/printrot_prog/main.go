package main

import (
	"fmt"
	"os"
)

func printAnswer(letter rune) {
	answer := []rune{}
	for i := letter; i <= 'z'; i++ {
		answer = append(answer, i)
	}
	for j := 'a'; j < letter; j++ {
		answer = append(answer, j)
	}
	fmt.Println(string(answer))
}

func main() {
	if len(os.Args) == 2 {
		letter := os.Args[1]
		if len(letter) == 1 {
			letterConverted := []rune(letter)
			if letterConverted[0] >= 'a' && letterConverted[0] <= 'z' {
				printAnswer(letterConverted[0])
			}
		} else {
			fmt.Println()
		}
	} else {
		fmt.Println()
	}
}
