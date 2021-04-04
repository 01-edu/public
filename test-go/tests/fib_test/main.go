package main

import (
	student "student"

	"github.com/01-edu/public/test-go/lib"
)

func fib(n int) int {
	if n <= 0 {
		return 0
	}
	t1 := 0
	t2 := 1
	for i := 2; i <= n; i++ {
		t1 += t2
		t1, t2 = t2, t1
	}
	return t2
}

func main() {
	table := append(lib.MultRandIntBetween(-100, 150),
		20,
		0,
		9,
		2,
	)
	for _, arg := range table {
		lib.Challenge("Fib", student.Fib, fib, arg)
	}
}
