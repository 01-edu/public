package main

import (
	"sort"

	student "student"

	"github.com/01-edu/public/test-go/lib"
)

func max(a []int) int {
	if len(a) == 0 {
		return 0
	}
	sort.Ints(a)
	return a[len(a)-1]
}

func main() {
	args := []int{lib.RandInt()}
	limit := lib.RandIntBetween(20, 50)
	for i := 0; i < limit; i++ {
		args = append(args, lib.RandInt())
	}

	lib.Challenge("Max", student.Max, max, args)
}
