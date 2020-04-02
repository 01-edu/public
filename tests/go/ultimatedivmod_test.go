package student_test

import (
	"testing"

	student "./student"
	"github.com/01-edu/z01"
)

func TestUltimateDivMod(t *testing.T) {
	i := 0
	for i < z01.SliceLen {
		a := z01.RandInt()
		b := z01.RandInt()
		aCopy := a
		bCopy := b
		div := a / b
		mod := a % b
		student.UltimateDivMod(&a, &b)
		if a != div {
			t.Errorf("DivMod(%d, %d), a == %d instead of %d", aCopy, bCopy, a, div)
		}
		if b != mod {
			t.Errorf("DivMod(%d, %d), b == %d instead of %d", aCopy, bCopy, b, mod)
		}
		i++
	}
}
