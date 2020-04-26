package main

import (
	"reflect"
	"strings"

	"github.com/01-edu/z01"

	solutions "./solutions"
	student "./student"
)

func main() {
	table := [][]string{{"a", "A", "1", "b", "B", "2", "c", "C", "3"}}

	table = append(table, z01.MultMultRandWords()...)

	for _, org := range table {
		// copy for using the solution function
		cp_sol := make([]string, len(org))
		// copy for using the student function
		cp_stu := make([]string, len(org))

		copy(cp_sol, org)
		copy(cp_stu, org)

		solutions.AdvancedSortWordArr(cp_sol, strings.Compare)
		student.AdvancedSortWordArr(cp_stu, strings.Compare)

		if !reflect.DeepEqual(cp_stu, cp_sol) {
			z01.Fatalf("%s(%v) == %v instead of %v\n",
				"AdvancedSortWordArr",
				org,
				cp_stu,
				cp_sol,
			)
		}
	}
}
