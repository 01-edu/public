package main

import (
	"fmt"
	"os"
)

func matchBrackets(exp string) bool {
	rn := []rune(exp)
	var opened []rune
	ptr := -1
	for _, c := range rn {
		if c == '(' || c == '[' || c == '{' {
			opened = append(opened, c)
			ptr++
		} else if c == ')' {
			if ptr < 0 || opened[ptr] != '(' {
				return false
			} else {
				opened = opened[:len(opened)-1]
				ptr--
			}
		} else if c == ']' {
			if ptr < 0 || opened[ptr] != '[' {
				return false
			} else {
				opened = opened[:len(opened)-1]
				ptr--
			}
		} else if c == '}' {
			if ptr < 0 || opened[ptr] != '{' {
				return false
			} else {
				opened = opened[:len(opened)-1]
				ptr--
			}
		}
	}
	return len(opened) == 0
}

func main() {
	if len(os.Args) > 1 {
		for _, v := range os.Args[1:] {
			if matchBrackets(v) {
				fmt.Println("OK")
			} else {
				fmt.Println("Error")
			}
		}
	} else {
		fmt.Println()
	}
}
