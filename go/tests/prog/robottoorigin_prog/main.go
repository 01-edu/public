package main

import (
	"fmt"
	"os"
)

func solve(s string) bool {
	x := 0
	y := 0

	for _, r := range s {
		switch r {
		case 'L':
			x--
		case 'R':
			x++
		case 'U':
			y++
		case 'D':
			y--
		}
	}
	return x == 0 && y == 0
}

func main() {
	if len(os.Args) == 2 {
		fmt.Println(solve(os.Args[1]))
	}
}
