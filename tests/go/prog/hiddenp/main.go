package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	first := os.Args[1]
	second := os.Args[2]

	if first == "" {
		fmt.Println("1")
		return
	}
	if second == "" {
		fmt.Println("0")
		return
	}

	i := 0
	j := 0
	count := 0

	firstA := []rune(first)
	secondA := []rune(second)

	for i < len(first) {
		for j < len(second) {
			if firstA[i] == secondA[j] {
				count++
				i++
			}
			if i == len(first) {
				break
			}
			j++
		}
		i++
	}

	if count == len(first) {
		fmt.Println("1")
	} else {
		fmt.Println("0")
	}
}
