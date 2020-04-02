package student_test

import (
	"reflect"
	"testing"

	"github.com/01-edu/z01"

	solutions "./solutions"
	student "./student"
)

func TestCompact(t *testing.T) {
	var arg [][]string

	for i := 0; i < 20; i++ {
		n := z01.RandIntBetween(5, 20) // random size of the slice

		orig := make([]string, n) //slice with the original value

		num_pos := z01.RandIntBetween(1, n-1) // random number of positions to be written

		for i := 0; i < num_pos; i++ {
			word := z01.RandWords()                // random string value
			rand_pos := z01.RandIntBetween(0, n-1) // random position in the slice
			orig[rand_pos] = word
		}
		arg = append(arg, orig)
	}

	arg = append(arg, []string{"hello", "", "hi", "", "salut", "", ""})

	for _, v := range arg {
		sli_sol := make([]string, len(arg)) // slice to apply the solution function
		sli_stu := make([]string, len(arg)) // slice to apply the student function

		copy(sli_sol, v)
		copy(sli_stu, v)

		sol_size := solutions.Compact(&sli_sol)
		stu_size := student.Compact(&sli_stu)

		if !reflect.DeepEqual(sli_stu, sli_sol) {
			t.Errorf("Produced slice: %v, instead of %v\n",
				sli_stu,
				sli_sol,
			)
		}

		if sol_size != stu_size {
			t.Errorf("%s(%v) == %v instead of %v\n",
				"Compact",
				v,
				sli_stu,
				sli_sol,
			)
		}
	}
}
