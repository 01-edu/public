package student_test

import (
	"reflect"
	"testing"

	"github.com/01-edu/z01"

	solutions "./solutions"
	student "./student"
)

func TestSortIntegerTable(t *testing.T) {
	i := 0
	for i < z01.SliceLen {
		table := z01.MultRandIntBetween(-100, 100)

		tableCopyBefore := make([]int, len(table))
		copy(tableCopyBefore, table)

		table2 := make([]int, len(table))
		copy(table2, table)

		student.SortIntegerTable(table)
		solutions.SortIntegerTable(table2)
		if !reflect.DeepEqual(table, table2) {
			t.Errorf("SortIntegerTable(%v), table == %v instead of %v ", tableCopyBefore, table, table2)
		}
		i++
	}
}
