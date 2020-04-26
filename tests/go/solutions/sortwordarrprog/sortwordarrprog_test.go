package main

import (
	"reflect"

	"github.com/01-edu/z01"

	solutions "../../solutions"
)

func main() {
	table := [][]string{{"a", "A", "1", "b", "B", "2", "c", "C", "3"}}

	for i := 0; i < 15; i++ {
		table = append(table, z01.MultRandWords())
	}

	for _, org := range table {
		// copy for using the solution function
		cp_sol := make([]string, len(org))
		// copy for using the student function
		cp_stu := make([]string, len(org))

		copy(cp_sol, org)
		copy(cp_stu, org)

		solutions.SortWordArr(cp_sol)
		SortWordArr(cp_stu)

		if !reflect.DeepEqual(cp_stu, cp_sol) {
			z01.Fatalf("%s(%v) == %v instead of %v\n",
				"SortWordArr",
				org,
				cp_stu,
				cp_sol,
			)
		}
	}
}
