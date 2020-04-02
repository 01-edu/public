package student_test

import (
	"testing"

	solutions "./solutions"
	student "./student"
	"github.com/01-edu/z01"
)

func TestRaid1e(t *testing.T) {
	// testing examples of subjects
	table := []int{
		5, 3,
		5, 1,
		1, 1,
		1, 5,
	}

	// testing special cases and one valid random case.
	table = append(table,
		0, 0,
		-1, 6,
		6, -1,
		z01.RandIntBetween(1, 20), z01.RandIntBetween(1, 20),
	)

	//Tests all possibilities including 0 0, -x y, x -y
	for i := 0; i < len(table); i = i + 2 {
		if i != len(table)-1 {
			z01.Challenge(t, solutions.Raid1e, student.Raid1e, table[i], table[i+1])
		}
	}
}
