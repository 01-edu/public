package main

import (
	"github.com/01-edu/z01"

	student "./student"
)

func main() {
	n := 0
	student.PointOne(&n)
	if n != 1 {
		z01.Fatalf("PointOne(&n), n == %d instead of 1\n", n)
	}
}
