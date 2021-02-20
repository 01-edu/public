package main

import (
	"math/rand"

	student "student"

	"github.com/01-edu/public/go/tests/lib"
)

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func main() {
	for i := 0; i < 20; i++ {
		token := rand.Perm(20)
		target := rand.Intn(30)
		lib.Challenge("TwoSum", student.TwoSum, twoSum, token, target)
	}
}
