package main

import (
	"testing"

	solutions "../../solutions"
	"github.com/01-edu/z01"
)

func TestFindPrevPrime(t *testing.T) {
	array := []int{5, 4, 1}
	for i := 0; i < 7; i++ {
		array = append(array, z01.RandIntBetween(0, 99999))
	}

	for i := 0; i < len(array); i++ {
		z01.Challenge(t, FindPrevPrime, solutions.FindPrevPrime, array[i])
	}
}
