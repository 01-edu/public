package main

import (
	student "student"

	"github.com/01-edu/public/test-go/lib"
)

func main() {
	i := 0
	for i < lib.SliceLen {
		a := lib.RandInt()
		b := lib.RandInt()
		aCopy := a
		bCopy := b
		div := a / b
		mod := a % b
		student.UltimateDivMod(&a, &b)
		if a != div {
			lib.Fatalf("DivMod(%d, %d), a == %d instead of %d", aCopy, bCopy, a, div)
		}
		if b != mod {
			lib.Fatalf("DivMod(%d, %d), b == %d instead of %d", aCopy, bCopy, b, mod)
		}
		i++
	}
}
