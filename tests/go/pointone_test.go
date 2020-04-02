package student_test

import (
	"testing"

	student "./student"
)

func TestPointOne(t *testing.T) {
	n := 0
	student.PointOne(&n)
	if n != 1 {
		t.Errorf("PointOne(&n), n == %d instead of 1", n)
	}
}
