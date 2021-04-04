package main

import (
	student "student"

	"github.com/01-edu/public/test-go/lib"
)

func priorPrime(x int) int {
	ans := 0
	ok := 0
	for i := 2; i < x; i++ {
		ok = 1
		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				ok = 0
			}
		}
		if ok == 1 {
			ans += i
		}
	}
	return ans
}

func main() {
	table := append(lib.MultRandIntBetween(0, 1000),
		50,
		13,
		10,
		0,
		1,
		2,
	)
	for _, arg := range table {
		lib.Challenge("PriorPrime", student.PriorPrime, priorPrime, arg)
	}
}
