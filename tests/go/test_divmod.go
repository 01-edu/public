package main

import (
	"github.com/01-edu/z01"

	student "./student"
)

func main() {
	i := 0
	for i < z01.SliceLen {
		a := z01.RandInt()
		b := z01.RandInt()
		var div int
		var mod int
		student.DivMod(a, b, &div, &mod)
		if div != a/b {
			z01.Fatalf("DivMod(%d, %d, &div, &mod), div == %d instead of %d", a, b, div, a/b)
		}
		if mod != a%b {
			z01.Fatalf("DivMod(%d, %d, &div, &mod), mod == %d instead of %d", a, b, mod, a%b)
		}
		i++
	}
}
