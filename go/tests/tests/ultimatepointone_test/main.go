package main

import (
	student "student"

	"github.com/01-edu/public/go/tests/lib"
)

func main() {
	a := 0
	b := &a
	n := &b
	student.UltimatePointOne(&n)
	if a != 1 {
		lib.Fatalf("UltimatePointOne(&n), a == %d instead of 1", a)
	}
}
