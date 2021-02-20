package main

import (
	"reflect"

	student "student"

	"github.com/01-edu/public/go/tests/lib"
)

func compact(slice *[]string) int {
	count := 0
	var compacted []string
	for _, v := range *slice {
		if v != "" {
			count++
			compacted = append(compacted, v)
		}
	}
	*slice = compacted
	return count
}

func main() {
	arg := [][]string{{"hello", "", "hi", "", "salut", "", ""}}

	for i := 0; i < 20; i++ {
		n := lib.RandIntBetween(5, 20)

		orig := make([]string, n)

		numPos := lib.RandIntBetween(1, n-1)

		for i := 0; i < numPos; i++ {
			word := lib.RandWords()
			randPos := lib.RandIntBetween(0, n-1)
			orig[randPos] = word
		}
		arg = append(arg, orig)
	}

	for _, v := range arg {
		solSlice := make([]string, len(arg))
		stuSlice := make([]string, len(arg))

		copy(solSlice, v)
		copy(stuSlice, v)

		solSize := compact(&solSlice)
		stuSize := student.Compact(&stuSlice)

		if !reflect.DeepEqual(stuSlice, solSlice) {
			lib.Fatalf("Produced slice: %v, instead of %v\n",
				stuSlice,
				solSlice,
			)
		}

		if solSize != stuSize {
			lib.Fatalf("%s(%v) == %v instead of %v\n",
				"Compact",
				v,
				stuSlice,
				solSlice,
			)
		}
	}
}
