package main

import (
	"piscine"
)

func main() {
	mul := func(acc int, cur int) int {
		return acc * cur
	}
	sum := func(acc int, cur int) int {
		return acc + cur
	}
	div := func(acc int, cur int) int {
		return acc / cur
	}
	as := []int{500, 2}
	piscine.ReduceInt(as, mul)
	piscine.ReduceInt(as, sum)
	piscine.ReduceInt(as, div)
}
