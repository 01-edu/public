package main

import (
	student "student"

	"github.com/01-edu/public/test-go/lib"
)

// Returns the element of the slice that doesn't have a correspondant pair
func unmatch(elems []int) int {
	var quant int
	for _, el := range elems {
		quant = 0
		for _, v := range elems {
			if v == el {
				quant++
			}
		}
		if quant%2 != 0 {
			return el
		}
	}
	return -1
}

func main() {
	i1 := lib.RandIntBetween(-100, 100)
	i2 := lib.RandIntBetween(-1000, 1000)
	i3 := lib.RandIntBetween(-10, 10)
	args := [][]int{
		{1, 1, 2, 3, 4, 3, 4},
		{1, 1, 2, 4, 3, 4, 2, 3, 4},
		{1, 2, 1, 1, 4, 5, 5, 4, 1, 7},
		{1, 2, 3, 4, 5, 6, 7, 8},
		{0, 20, 91, 23, 10, 34},
		{1, 1, 2, 2, 3, 4, 3, 4, 5, 5, 8, 9, 8, 9},
		{i1, i2, i1, i2, i1 + i3, i1 + i3},
		{i1, i2, i1, i2, i1 + i3, i2 - i3},
	}
	for _, v := range args {
		lib.Challenge("Unmatch", student.Unmatch, unmatch, v)
	}
}
