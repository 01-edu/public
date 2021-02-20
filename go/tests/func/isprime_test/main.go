package main

import (
	student "student"

	"github.com/01-edu/public/go/tests/lib"
	"github.com/01-edu/public/go/tests/lib/is"
)

func main() {
	table := append(
		lib.MultRandIntBetween(-1000000, 1000000),
		0,
		1,
		2,
		3,
		4,
		5,
		6,
		7,
		8,
		9,
		10,
		11,
		12,
		100,
		1000000086,
		1000000087,
	)
	for _, arg := range table {
		lib.Challenge("IsPrime", student.IsPrime, is.Prime, arg)
	}
}
