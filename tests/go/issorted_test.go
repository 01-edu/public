package student_test

import (
	"sort"
	"testing"

	"github.com/01-edu/z01"

	solutions "./solutions"
	student "./student"
)

func TestIsSorted(t *testing.T) {
	functionsArray := []func(int, int) int{solutions.IsSortedByDiff, solutions.IsSortedBy1, solutions.IsSortedBy10}

	type node struct {
		f   func(int, int) int
		arr []int
	}

	table := []node{}

	//5 unordered slices
	for i := 0; i < 5; i++ {

		functionSelected := functionsArray[z01.RandIntBetween(0, len(functionsArray)-1)]
		val := node{
			f:   functionSelected,
			arr: z01.MultRandIntBetween(-1000000, 1000000),
		}
		table = append(table, val)

	}

	//5 slices ordered in ascending order
	for i := 0; i < 5; i++ {

		functionSelected := functionsArray[z01.RandIntBetween(0, len(functionsArray)-1)]
		orderedArr := z01.MultRandIntBetween(-1000000, 1000000)
		sort.Ints(orderedArr)

		val := node{
			f:   functionSelected,
			arr: orderedArr,
		}
		table = append(table, val)

	}

	//5 slices ordered in descending order
	for i := 0; i < 5; i++ {

		functionSelected := functionsArray[z01.RandIntBetween(0, len(functionsArray)-1)]
		reverseArr := z01.MultRandIntBetween(-1000000, 1000000)
		sort.Sort(sort.Reverse(sort.IntSlice(reverseArr)))
		val := node{
			f:   functionSelected,
			arr: reverseArr,
		}
		table = append(table, val)

	}

	table = append(table, node{
		f:   solutions.IsSortedByDiff,
		arr: []int{1, 2, 3, 4, 5, 6},
	})

	table = append(table, node{
		f:   solutions.IsSortedByDiff,
		arr: []int{6, 5, 4, 3, 2, 1},
	})

	table = append(table, node{
		f:   solutions.IsSortedByDiff,
		arr: []int{0, 0, 0, 0, 0, 0, 0},
	})

	table = append(table, node{
		f:   solutions.IsSortedByDiff,
		arr: []int{0},
	})

	for _, arg := range table {
		z01.Challenge(t, student.IsSorted, solutions.IsSorted, arg.f, arg.arr)
	}
}
