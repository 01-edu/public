package main

import "./student"

func UltimateDivMod(a, b *int) {
	temp := *a
	*a = *a / *b
	*b = temp % *b
}

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
