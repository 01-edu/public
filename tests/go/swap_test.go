package main

import (
	"github.com/01-edu/z01"

	student "./student"
)

func main() {
	i := 0
	for i < 30 {
		a := z01.RandInt()
		b := z01.RandInt()
		aCopy := a
		bCopy := b
		student.Swap(&a, &b)
		if a != bCopy {
			z01.Fatalf("Swap(%d, %d), a == %d instead of %d", aCopy, bCopy, a, bCopy)
		}
		if b != aCopy {
			z01.Fatalf("Swap(%d, %d), b == %d instead of %d", aCopy, bCopy, b, aCopy)
		}
		i++
	}
}
