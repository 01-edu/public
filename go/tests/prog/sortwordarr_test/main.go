package main

import (
	"reflect"
	"sort"

	student "student"

	"github.com/01-edu/public/go/tests/lib"
)

func main() {
	table := [][]string{{"a", "A", "1", "b", "B", "2", "c", "C", "3"}}

	for i := 0; i < 15; i++ {
		table = append(table, lib.MultRandWords())
	}

	for _, org := range table {
		// copy for using the solution function
		copySol := make([]string, len(org))
		// copy for using the student function
		copyStu := make([]string, len(org))

		copy(copySol, org)
		copy(copyStu, org)

		sort.Strings(copySol)
		student.SortWordArr(copyStu)

		if !reflect.DeepEqual(copyStu, copySol) {
			lib.Fatalf("%s(%v) == %v instead of %v\n",
				"SortWordArr",
				org,
				copyStu,
				copySol,
			)
		}
	}
}
