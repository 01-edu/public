package main

import (
    "math/rand"
	"testing"

	"github.com/01-edu/z01"
	solutions "../../solutions"
)

func TestTwoSum(t *testing.T) {
    for i := 0; i < 20; i++ {
    	token := rand.Perm(20)
	target := rand.Intn(30)
    	z01.Challenge(t, TwoSum, solutions.TwoSum, token, target)
    }
}
