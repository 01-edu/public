package main

import (
	"./student"
	"github.com/01-edu/public/go/lib"
)

func fib(n int) int {
	if n <= 0 {
		return 0
	}
	t1 := 0
	t2 := 1
	for i := 2; i <= n; i++ {
		t1 += t2
		tmp := t1
		t1 = t2
		t2 = tmp
	}
	return t2
}

func main() {
	table := []int{
		20,
		0,
		9,
		2,
	}
	table = append(table, lib.MultRandIntBetween(-100, 150)...)
	for _, arg := range table {
		lib.Challenge("Fib", student.Fib, fib, arg)
	}
}
