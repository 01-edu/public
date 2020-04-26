package main

import (
	"math/rand"

	"github.com/01-edu/z01"

	correct "./correct"
	student "./student"
)

func main() {
	for i := 0; i < 20; i++ {
		token := rand.Perm(20)
		target := rand.Intn(30)
		z01.Challenge("TwoSum", student.TwoSum, correct.TwoSum, token, target)
	}
}
