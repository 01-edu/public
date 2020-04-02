package main

import (
	"testing"

	"github.com/01-edu/z01"

	solutions "../../solutions"
)

func TestMaxProg(t *testing.T) {
	args := []int{z01.RandInt()}
	limit := z01.RandIntBetween(20, 50)
	for i := 0; i < limit; i++ {
		args = append(args, z01.RandInt())
	}

	z01.Challenge(t, Max, solutions.Max, args)
}
