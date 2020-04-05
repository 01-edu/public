package main

import (
	"math/rand"
	"testing"

	solutions "../../solutions"
	"github.com/01-edu/z01"
)

func TestTwoSum(t *testing.T) {
	for i := 0; i < 20; i++ {
		token := rand.Perm(20)
		target := rand.Intn(30)
		z01.Challenge(t, TwoSum, solutions.TwoSum, token, target)
	}
}
