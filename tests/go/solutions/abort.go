package solutions

import (
	"sort"
)

//Receives 5 ints and returns the number in the middle
func Abort(a, b, c, d, e int) int {
	arg := []int{a, b, c, d, e}
	sort.Sort(sort.IntSlice(arg))
	return arg[2]
}
