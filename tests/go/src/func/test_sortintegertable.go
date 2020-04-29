package main

import (
	"reflect"

	"../lib"
	"./correct"
	"./student"
)

func main() {
	i := 0
	for i < lib.SliceLen {
		table1 := lib.MultRandIntBetween(-100, 100)

		tableCopyBefore := make([]int, len(table1))
		copy(tableCopyBefore, table1)

		table2 := make([]int, len(table1))
		copy(table2, table1)

		student.SortIntegerTable(table1)
		correct.SortIntegerTable(table2)
		if !reflect.DeepEqual(table1, table2) {
			lib.Fatalf("SortIntegerTable(%v), table1 == %v instead of %v ", tableCopyBefore, table1, table2)
		}
		i++
	}
}
