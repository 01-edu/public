package student_test

import (
	"testing"

	student "./student"
)

func TestUltimatePointOne(t *testing.T) {
	a := 0
	b := &a
	n := &b
	student.UltimatePointOne(&n)
	if a != 1 {
		t.Errorf("UltimatePointOne(&n), a == %d instead of 1", a)
	}
}
