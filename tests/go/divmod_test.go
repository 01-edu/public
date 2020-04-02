package student_test

import (
	"testing"

	"github.com/01-edu/z01"

	student "./student"
)

func TestDivMod(t *testing.T) {
	i := 0
	for i < z01.SliceLen {
		a := z01.RandInt()
		b := z01.RandInt()
		var div int
		var mod int
		student.DivMod(a, b, &div, &mod)
		if div != a/b {
			t.Errorf("DivMod(%d, %d, &div, &mod), div == %d instead of %d", a, b, div, a/b)
		}
		if mod != a%b {
			t.Errorf("DivMod(%d, %d, &div, &mod), mod == %d instead of %d", a, b, mod, a%b)
		}
		i++
	}
}
