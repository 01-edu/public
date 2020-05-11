package main

import "./student"

func UltimatePointOne(n ***int) {
	***n = 1
}

func main() {
	a := 0
	b := &a
	n := &b
	student.UltimatePointOne(&n)
	if a != 1 {
		lib.Fatalf("UltimatePointOne(&n), a == %d instead of 1", a)
	}
}
