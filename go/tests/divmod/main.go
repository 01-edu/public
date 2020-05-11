package main

import (
	"./student"
	"github.com/01-edu/public/go/lib"
)

func DivMod(a, b int, div, mod *int) {
	*div = a / b
	*mod = a % b
}

func main() {
	i := 0
	for i < lib.SliceLen {
		a := lib.RandInt()
		b := lib.RandInt()
		var div int
		var mod int
		student.DivMod(a, b, &div, &mod)
		if div != a/b {
			lib.Fatalf("DivMod(%d, %d, &div, &mod), div == %d instead of %d", a, b, div, a/b)
		}
		if mod != a%b {
			lib.Fatalf("DivMod(%d, %d, &div, &mod), mod == %d instead of %d", a, b, mod, a%b)
		}
		i++
	}
}
