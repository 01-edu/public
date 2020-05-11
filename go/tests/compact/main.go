package main

import (
	"reflect"

	"./student"
	"github.com/01-edu/public/go/lib"
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
		n := lib.RandIntBetween(5, 20) // random size of the slice

		orig := make([]string, n) // slice with the original value

		num_pos := lib.RandIntBetween(1, n-1) // random number of positions to be written

		for i := 0; i < num_pos; i++ {
			word := lib.RandWords()                // random string value
			rand_pos := lib.RandIntBetween(0, n-1) // random position in the slice
			orig[rand_pos] = word
		}
		arg = append(arg, orig)
	}

	for _, v := range arg {
		sli_sol := make([]string, len(arg)) // slice to apply the solution function
		sli_stu := make([]string, len(arg)) // slice to apply the student function

		copy(sli_sol, v)
		copy(sli_stu, v)

		sol_size := compact(&sli_sol)
		stu_size := student.Compact(&sli_stu)

		if !reflect.DeepEqual(sli_stu, sli_sol) {
			lib.Fatalf("Produced slice: %v, instead of %v\n",
				sli_stu,
				sli_sol,
			)
		}

		if sol_size != stu_size {
			lib.Fatalf("%s(%v) == %v instead of %v\n",
				"Compact",
				v,
				sli_stu,
				sli_sol,
			)
		}
	}
}
