package main

import (
	"github.com/01-edu/z01"
	"os"
	"strconv"
)

func printLine(str string) {
	for _, c := range str {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}

func solve(x, y int) {
	for i := 0; i < x; i++ {
		line := ""
		for j := 0; j < y; j++ {
			if (i%2 == 0) && (j%2 == 0) {
				line += "#"
			} else if (i%2 == 0) && (j%2 == 1) {
				line += " "
			} else if (i%2 == 1) && (j%2 == 1) {
				line += "#"
			} else {
				line += " "
			}
		}
		printLine(line)
	}
}

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		printLine("Error")
		return
	}
	x, _ := strconv.Atoi(args[1])
	y, _ := strconv.Atoi(args[0])
	if x <= 0 || y <= 0 {
		printLine("Error")
		return
	}
	solve(x, y)
}
