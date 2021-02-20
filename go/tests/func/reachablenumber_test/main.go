package main

import (
	student "student"

	"github.com/01-edu/public/go/tests/lib"
)

func reachableNumber(n int) int {
	cnt := 0
	for n > 0 {
		cnt++
		if n < 10 {
			cnt += 8
			break
		} else {
			n++
		}
		for n%10 == 0 {
			n /= 10
		}
	}
	return cnt
}

func main() {
	table := []int{
		20,
		1,
		9,
		2,
	}
	for i := 0; i < 25; i++ {
		table = append(table, lib.MultRandIntBetween(1, 877)...)
	}
	for _, arg := range table {
		lib.Challenge("ReachableNumber", student.ReachableNumber, reachableNumber, arg)
	}
}
