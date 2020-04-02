package solutions

import (
	"sort"
)

func SortWordArr(array []string) {
	sort.Sort(sort.StringSlice(array))
}
