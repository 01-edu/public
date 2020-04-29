package main

import (
	"sort"

	"../lib"
	"./correct"
	"./student"
)

func isSortedBy1(a, b int) int {
	if a-b < 0 {
		return -1
	}
	if a-b > 0 {
		return 1
	}
	return 0
}

func isSortedBy10(a, b int) int {
	if a-b < 0 {
		return -10
	}
	if a-b > 0 {
		return 10
	}
	return 0
}

func isSortedByDiff(a, b int) int {
	return a - b
}

func main() {
	functions := []func(int, int) int{isSortedByDiff, isSortedBy1, isSortedBy10}

	type node struct {
		f func(int, int) int
		a []int
	}

	table := []node{}

	// 5 unordered slices
	for i := 0; i < 5; i++ {
		function := functions[lib.RandIntBetween(0, len(functions)-1)]
		val := node{
			f: function,
			a: lib.MultRandIntBetween(-1000000, 1000000),
		}
		table = append(table, val)
	}

	// 5 slices ordered in ascending order
	for i := 0; i < 5; i++ {
		function := functions[lib.RandIntBetween(0, len(functions)-1)]
		ordered := lib.MultRandIntBetween(-1000000, 1000000)
		sort.Ints(ordered)

		val := node{
			f: function,
			a: ordered,
		}
		table = append(table, val)
	}

	// 5 slices ordered in descending order
	for i := 0; i < 5; i++ {
		function := functions[lib.RandIntBetween(0, len(functions)-1)]
		reversed := lib.MultRandIntBetween(-1000000, 1000000)
		sort.Sort(sort.Reverse(sort.IntSlice(reversed)))
		val := node{
			f: function,
			a: reversed,
		}
		table = append(table, val)
	}

	table = append(table, node{
		f: correct.IsSortedByDiff,
		a: []int{1, 2, 3, 4, 5, 6},
	})

	table = append(table, node{
		f: correct.IsSortedByDiff,
		a: []int{6, 5, 4, 3, 2, 1},
	})

	table = append(table, node{
		f: correct.IsSortedByDiff,
		a: []int{0, 0, 0, 0, 0, 0, 0},
	})

	table = append(table, node{
		f: correct.IsSortedByDiff,
		a: []int{0},
	})

	for _, arg := range table {
		lib.Challenge("IsSorted", student.IsSorted, correct.IsSorted, arg.f, arg.a)
	}
}
