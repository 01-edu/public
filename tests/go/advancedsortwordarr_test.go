package student_test

import (
	"reflect"
	"testing"

	"github.com/01-edu/z01"

	solutions "./solutions"
	student "./student"
)

func TestAdvancedSortWordArr(t *testing.T) {
	var table [][]string

	for i := 0; i < 10; i++ {
		table = append(table, z01.MultRandWords())
	}

	table = append(table, []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"})

	for _, org := range table {
		//copy for using the solution function
		cp_sol := make([]string, len(org))
		//copy for using the student function
		cp_stu := make([]string, len(org))

		copy(cp_sol, org)
		copy(cp_stu, org)

		solutions.AdvancedSortWordArr(cp_sol, solutions.CompArray)
		student.AdvancedSortWordArr(cp_stu, solutions.CompArray)

		if !reflect.DeepEqual(cp_stu, cp_sol) {
			t.Errorf("%s(%v) == %v instead of %v\n",
				"AdvancedSortWordArr",
				org,
				cp_stu,
				cp_sol,
			)
		}
	}

}
