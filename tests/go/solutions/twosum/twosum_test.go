package main

import (
	"math/rand"

	solutions "../../solutions"
	"github.com/01-edu/z01"
)

func main() {
	for i := 0; i < 20; i++ {
		token := rand.Perm(20)
		target := rand.Intn(30)
		z01.Challenge("TwoSum", TwoSum, solutions.TwoSum, token, target)
	}
}
