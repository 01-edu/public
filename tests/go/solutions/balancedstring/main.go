package main

import (
	"fmt"
	"os"
)

func solve(str string) int {
	var count int = 0
	var countC, countD int
	for i := 0; i < len(str); i++ {
		if str[i] == 'C' {
			countC++
		} else if str[i] == 'D' {
			countD++
		}
		if countC == countD {
			count++
			countC = 0
			countD = 0
		}
	}
	return count
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println()
		return
	}
	result := solve(args[0])
	fmt.Println(result)
}
