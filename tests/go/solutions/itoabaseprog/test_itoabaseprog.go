package main

import (
	"github.com/01-edu/z01"

	solutions "../../solutions"
)

func main() {
	for i := 0; i < 30; i++ {
		value := z01.RandIntBetween(-1000000, 1000000)
		base := z01.RandIntBetween(2, 16)
		z01.Challenge("ItoaBaseProg", ItoaBase, solutions.ItoaBase, value, base)
	}
	for i := 0; i < 5; i++ {
		base := z01.RandIntBetween(2, 16)
		z01.Challenge("ItoaBaseProg", ItoaBase, solutions.ItoaBase, z01.MaxInt, base)
		z01.Challenge("ItoaBaseProg", ItoaBase, solutions.ItoaBase, z01.MinInt, base)
	}
}
