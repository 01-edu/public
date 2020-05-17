package main

import (
	"fmt"
	"os"
)

func solve(s string) int {
	var count, countC, countD int
	for _, r := range s {
		if r == 'C' {
			countC++
		} else if r == 'D' {
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
	if len(args) == 1 {
		result := solve(args[0])
		fmt.Println(result)
	}
}
