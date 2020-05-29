package main

import (
	"github.com/01-edu/z01"

	"./base"

	correct "./correct"
	student "./student"
)

func main() {
	type node struct {
		nbr      string
		baseFrom string
		baseTo   string
	}
	table := []node{}

	for i := 0; i < 30; i++ {
		validBaseToInput1 := base.Valid()
		validBaseToInput2 := base.Valid()
		str := base.ConvertNbr(z01.RandIntBetween(0, 1000000), validBaseToInput1)
		val := node{
			nbr:      str,
			baseFrom: validBaseToInput1,
			baseTo:   validBaseToInput2,
		}
		table = append(table, val)
	}

	table = append(table, node{
		nbr:      "101011",
		baseFrom: "01",
		baseTo:   "0123456789"},
	)

	for _, arg := range table {
		z01.Challenge("ConvertBase", student.ConvertBase, correct.ConvertBase, arg.nbr, arg.baseFrom, arg.baseTo)
	}
}

// TODO: fix base exercises
