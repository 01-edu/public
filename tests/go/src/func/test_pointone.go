package main

import (
	"../lib"
	"./student"
)

func main() {
	n := 0
	student.PointOne(&n)
	if n != 1 {
		lib.Fatalf("PointOne(&n), n == %d instead of 1\n", n)
	}
}
