package main

import (
	"reflect"
	"sort"
	"strings"

	student "student"

	"github.com/01-edu/public/go/tests/lib"
)

func advancedSortWordArr(a []string, f func(a, b string) int) {
	sort.Slice(a, func(i, j int) bool {
		return f(a[i], a[j]) < 0
	})
}

func main() {
	table := [][]string{{"a", "A", "1", "b", "B", "2", "c", "C", "3"}}

	table = append(table, lib.MultRandWords())

	for _, org := range table {
		// copy for using the solution function
		cp_sol := make([]string, len(org))
		// copy for using the student function
		cp_stu := make([]string, len(org))

		copy(cp_sol, org)
		copy(cp_stu, org)

		advancedSortWordArr(cp_sol, strings.Compare)
		student.AdvancedSortWordArr(cp_stu, strings.Compare)

		if !reflect.DeepEqual(cp_stu, cp_sol) {
			lib.Fatalf("%s(%v) == %v instead of %v\n",
				"AdvancedSortWordArr",
				org,
				cp_stu,
				cp_sol,
			)
		}
	}
}
