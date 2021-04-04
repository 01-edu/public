package main

import (
	"fmt"
	"sort"

	student "student"

	"github.com/01-edu/public/go/tests/lib"
)

func intToDigits(n int) (digits []int) {
	for n > 0 {
		if n == 0 {
			digits = append(digits, 0)
		} else {
			digits = append(digits, n%10)
		}
		n /= 10
	}
	return
}

func printNbrInOrder(n int) {
	if n == 0 {
		fmt.Print("0")
		return
	}
	digits := intToDigits(n)
	sort.Ints(digits)
	for _, i := range digits {
		fmt.Printf("%c", rune(i)+'0')
	}
}

func main() {
	table := append(
		lib.MultRandIntBetween(0, lib.MaxInt),
		lib.MaxInt,
		321,
		321321,
		0,
	)
	for _, arg := range table {
		lib.Challenge("PrintNbrInOrder", student.PrintNbrInOrder, printNbrInOrder, arg)
	}
}
