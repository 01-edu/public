package main

import "fmt"

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
	if len(args) == 2 {
		fmt.Println(solve(args[1]))
	}
}
