package main

import (
	"../lib"
	"./correct"
	"./student"
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
	)
	for _, arg := range table {
		lib.Challenge("Sqrt", student.Sqrt, correct.Sqrt, arg)
	}
}
