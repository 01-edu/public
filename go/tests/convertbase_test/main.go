package main

import (
	student "student"

	"github.com/01-edu/public/go/tests/lib"
	"github.com/01-edu/public/go/tests/lib/base"
)

func convertNbrBase(n int, base string) string {
	var result string
	length := len(base)

	for n >= length {
		result = string(base[(n%length)]) + result
		n /= length
	}
	result = string(base[n]) + result

	return result
}

func convertBase(nbr, baseFrom, baseTo string) string {
	resultIntermediary := base.Atoi(nbr, baseFrom)

	resultFinal := convertNbrBase(resultIntermediary, baseTo)

	return resultFinal
}

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
		str := base.ConvertNbr(lib.RandIntBetween(0, 1000000), validBaseToInput1)
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
		baseTo:   "0123456789",
	})

	for _, arg := range table {
		lib.Challenge("ConvertBase", student.ConvertBase, convertBase, arg.nbr, arg.baseFrom, arg.baseTo)
	}
}

// TODO: fix base exercises
