package main

import (
	"github.com/01-edu/z01"
	"os"
)

func solve(str string) bool {
	x := 0
	y := 0

	for _, c := range str {
		if c == 'L' {
			x--
		} else if c == 'R' {
			x++
		} else if c == 'U' {
			y++
		} else if c == 'D' {
			y--
		}
	}
	if x == 0 && y == 0 {
		return true
	}
	return false
}

func print(str string) {
	for _, v := range str {
		z01.PrintRune(v)
	}
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		z01.PrintRune('\n')
		return
	}
	result := solve(args[0])
	if result {
		print("true\n")
	} else {
		print("false\n")
	}
}
