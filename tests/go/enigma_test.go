package student_test

import (
	solutions "./solutions"
	student "./student"
	"github.com/01-edu/z01"
	"testing"
)

func TestEnigma(t *testing.T) {
	args := []int{z01.RandIntBetween(2, 20)}
	for i := 0; i < 3; i++ {
		args = append(args, z01.RandIntBetween(2, 20))
	}

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
	solutions.Decript(a, b, c, d)

	if aval != ***a {
		t.Errorf("Expected ***a = %d instead of %d\n",
			aval,
			***a,
		)
	}
	if bval != *b {
		t.Errorf("Expected *b = %d instead of %d\n",
			bval,
			*b,
		)
	}
	if cval != *******c {
		t.Errorf("Expected *******c = %d instead of %d\n",
			cval,
			*******c,
		)
	}
	if dval != ****d {
		t.Errorf("Expected ****d = %d instead of %d\n",
			dval,
			****d,
		)
	}
}
