package main

import (
	student "student"

	"github.com/01-edu/public/go/tests/lib"
)

func isInteresting(n int) bool {
	s := 0
	for n > 0 {
		s += n % 10
		n /= 10
	}
	return s%7 == 0
}

func interestingNumber(n int) int {
	for {
		if isInteresting(n) {
			return n
		}
		n++
	}
}

func main() {
	table := []int{
		20,
		1,
		9,
		2,
	}
	table = append(table, lib.MultRandIntBetween(1, 1500)...)

	for _, arg := range table {
		lib.Challenge("InterestingNumber", student.InterestingNumber, interestingNumber, arg)
	}
}
