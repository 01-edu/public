package main

import (
	"testing"

	"github.com/01-edu/z01"
)

func TestSwapProg(t *testing.T) {
	i := 0
	for i < 30 {
		a := z01.RandInt()
		b := z01.RandInt()
		aCopy := a
		bCopy := b
		Swap(&a, &b)
		if a != bCopy {
			t.Errorf("Swap(%d, %d), a == %d instead of %d", aCopy, bCopy, a, bCopy)
		}
		if b != aCopy {
			t.Errorf("Swap(%d, %d), b == %d instead of %d", aCopy, bCopy, b, aCopy)
		}
		i++
	}
}
