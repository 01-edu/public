package main

import (
	"strconv"

	"github.com/01-edu/public/go/lib"
)

type node struct {
	A, B, C, D, E, F int
}

func main() {
	table := []node{}

	for i := 0; i < 25; i++ {
		a := lib.RandIntBetween(0, 1000)
		b := lib.RandIntBetween(0, 1000)
		c := lib.RandIntBetween(0, 1000)
		d := lib.RandIntBetween(0, 1000)
		e := lib.RandIntBetween(0, 1000)
		f := lib.RandIntBetween(0, 1000)
		table = append(table, node{a, b, c, d, e, f})
	}

	for _, arg := range table {
		a := strconv.Itoa(arg.A)
		b := strconv.Itoa(arg.B)
		c := strconv.Itoa(arg.C)
		d := strconv.Itoa(arg.D)
		e := strconv.Itoa(arg.E)
		f := strconv.Itoa(arg.F)
		lib.ChallengeMain("costumeprofit", a, b, c, d, e, f)
	}
}
