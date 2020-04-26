package solutions

import "sort"

func Max(a []int) int {
	if len(a) == 0 {
		return 0
	}
	sort.Ints(a)
	return a[len(a)-1]
}
