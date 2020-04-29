package main

import (
	"math/rand"

	"../lib"
	"./correct"
	"./student"
)

func main() {
	for i := 0; i < 20; i++ {
		token := rand.Perm(20)
		target := rand.Intn(30)
		lib.Challenge("TwoSum", student.TwoSum, correct.TwoSum, token, target)
	}
}
