package main

import (
	"sort"

	student "student"

	"github.com/01-edu/public/go/tests/lib"
)

// Receives 5 ints and returns the number in the middle
func abort(a, b, c, d, e int) int {
	arg := []int{a, b, c, d, e}
	sort.Ints(arg)
	return arg[2]
}

func main() {
	arg := lib.MultRandInt()
	arg = append(arg, lib.RandInt())
	for i := 0; i < 15; i++ {
		lib.Challenge("Abort", student.Abort, abort, arg[0], arg[1], arg[2], arg[3], arg[4])
		arg = lib.MultRandInt()
		arg = append(arg, lib.RandInt())
	}
}
