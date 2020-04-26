package main

import (
	"../lib"
	"./base"
	"./correct"
	"./student"
)

func main() {
	type node struct {
		n    int
		base string
	}

	table := []node{}

	// 15 random pairs of ints with valid bases
	for i := 0; i < 15; i++ {
		validBaseToInput := base.Valid()
		val := node{
			n:    lib.RandIntBetween(-1000000, 1000000),
			base: validBaseToInput,
		}
		table = append(table, val)
	}

	// 15 random pairs of ints with invalid bases
	for i := 0; i < 15; i++ {
		invalidBaseToInput := base.Invalid()
		val := node{
			n:    lib.RandIntBetween(-1000000, 1000000),
			base: invalidBaseToInput,
		}
		table = append(table, val)
	}

	table = append(table,
		node{n: 125, base: "0123456789"},
		node{n: -125, base: "01"},
		node{n: 125, base: "0123456789ABCDEF"},
		node{n: -125, base: "choumi"},
		node{n: 125, base: "-ab"},
		node{n: lib.MinInt, base: "0123456789"},
	)
	for _, arg := range table {
		lib.Challenge("PrintNbrBase", student.PrintNbrBase, correct.PrintNbrBase, arg.n, arg.base)
	}
}

// TODO: fix base exercises
