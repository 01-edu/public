package main

import (
	"strings"

	student "student"

	"github.com/01-edu/public/go/tests/lib"
)

type node struct {
	big, little string
}

func main() {
	table := []node{}

	table = append(table,
		node{"aaaaaaa", "a"},
		node{"qwerty", "t"},
		node{"a", "b"},
	)

	for i := 0; i < 10; i++ {
		l := lib.RandIntBetween(5, 30)
		big := lib.RandStr(l, "qwertyuiopasdfghjklzxcvbnm")
		start := lib.RandIntBetween(0, l-1)
		end := lib.RandIntBetween(start+1, l-1)
		little := big[start:end]

		value := node{
			big:    big,
			little: little,
		}

		table = append(table, value)
	}

	for i := 0; i < 10; i++ {
		big := lib.RandStr(lib.RandIntBetween(5, 30), "qwertyuiopasdfghjklzxcvbnm")
		little := lib.RandStr(lib.RandIntBetween(1, 29), "qwertyuiopasdfghjklzxcvbnm")

		value := node{
			big:    big,
			little: little,
		}

		table = append(table, value)
	}

	for _, arg := range table {
		lib.Challenge("DoppelGanger", student.DoppelGanger, strings.LastIndex, arg.big, arg.little)
	}
}
