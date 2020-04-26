package main

import (
	student "./student"
	"github.com/01-edu/z01"
)

func main() {
	a := 0
	b := &a
	n := &b
	student.UltimatePointOne(&n)
	if a != 1 {
		z01.Fatalf("UltimatePointOne(&n), a == %d instead of 1", a)
	}
}
