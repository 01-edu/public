package main

import "./student"

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
	table := []int{
		50,
		13,
		10,
		0,
		1,
		2,
	}
	table = append(table, lib.MultRandIntBetween(0, 1000))
	for _, arg := range table {
		lib.Challenge("PriorPrime", student.PriorPrime, priorPrime, arg)
	}
}
