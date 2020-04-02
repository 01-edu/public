package main

import (
	"sort"
)

func Max(arr []int) int {
	sort.Sort(sort.IntSlice(arr))
	return arr[len(arr)-1]
}

func main() {

}
