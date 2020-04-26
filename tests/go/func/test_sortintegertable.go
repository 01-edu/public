package main

import (
	"reflect"

	"github.com/01-edu/z01"

	correct "./correct"
	student "./student"
)

func main() {
	i := 0
	for i < z01.SliceLen {
		table1 := z01.MultRandIntBetween(-100, 100)

		tableCopyBefore := make([]int, len(table1))
		copy(tableCopyBefore, table1)

		table2 := make([]int, len(table1))
		copy(table2, table1)

		student.SortIntegerTable(table1)
		correct.SortIntegerTable(table2)
		if !reflect.DeepEqual(table1, table2) {
			z01.Fatalf("SortIntegerTable(%v), table1 == %v instead of %v ", tableCopyBefore, table1, table2)
		}
		i++
	}
}
