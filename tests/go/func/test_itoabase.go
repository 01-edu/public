package main

import (
	"../lib"
	"./correct"
	"./student"
)

func main() {
	for i := 0; i < 30; i++ {
		value := lib.RandIntBetween(-1000000, 1000000)
		base := lib.RandIntBetween(2, 16)
		lib.Challenge("ItoaBase", student.ItoaBase, correct.ItoaBase, value, base)
	}
	for i := 0; i < 5; i++ {
		base := lib.RandIntBetween(2, 16)
		lib.Challenge("ItoaBase", student.ItoaBase, correct.ItoaBase, lib.MaxInt, base)
		lib.Challenge("ItoaBase", student.ItoaBase, correct.ItoaBase, lib.MinInt, base)
	}
}
