package main

import (
	"github.com/01-edu/z01"

	correct "./correct"
	student "./student"
)

func main() {
	args := append([]int, z01.MultRandIntBetween(2, 20)...)

	aval := args[0]
	x := args[0]
	y := &x
	z := &y
	a := &z

	bval := args[1]
	w := args[1]
	b := &w

	cval := args[2]
	u := args[2]
	e := &u
	f := &e
	g := &f
	h := &g
	i := &h
	j := &i
	c := &j

	dval := args[3]
	k := args[3]
	l := &k
	m := &l
	n := &m
	d := &n

	student.Enigma(a, b, c, d)
	correct.Decript(a, b, c, d)

	if aval != ***a {
		z01.Fatalf("Expected ***a = %d instead of %d\n",
			aval,
			***a,
		)
	}
	if bval != *b {
		z01.Fatalf("Expected *b = %d instead of %d\n",
			bval,
			*b,
		)
	}
	if cval != *******c {
		z01.Fatalf("Expected *******c = %d instead of %d\n",
			cval,
			*******c,
		)
	}
	if dval != ****d {
		z01.Fatalf("Expected ****d = %d instead of %d\n",
			dval,
			****d,
		)
	}
}

// TODO: remove all those pointers...
