package main

import (
	"github.com/01-edu/z01"

	correct "./correct"
)

func main() {
	args := []int{z01.RandInt()}
	limit := z01.RandIntBetween(20, 50)
	for i := 0; i < limit; i++ {
		args = append(args, z01.RandInt())
	}

	z01.Challenge("MaxProg", Max, correct.Max, args)
}
