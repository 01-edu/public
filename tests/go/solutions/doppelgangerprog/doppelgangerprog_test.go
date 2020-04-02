package main

import (
	"testing"

	solutions "../../solutions"
	"github.com/01-edu/z01"
)

type node struct {
	big, little string
}

func TestDoppelGanger(t *testing.T) {
	table := []node{}

	table = append(table,
		node{"aaaaaaa", "a"},
		node{"qwerty", "t"},
		node{"a", "b"},
	)

	for i := 0; i < 10; i++ {
		l := z01.RandIntBetween(5, 30)
		big := z01.RandStr(l, "qwertyuiopasdfghjklzxcvbnm")
		start := z01.RandIntBetween(0, l-1)
		end := z01.RandIntBetween(start+1, l-1)
		little := big[start:end]

		value := node{
			big:    big,
			little: little,
		}

		table = append(table, value)
	}

	for i := 0; i < 10; i++ {
		big := z01.RandStr(z01.RandIntBetween(5, 30), "qwertyuiopasdfghjklzxcvbnm")
		little := z01.RandStr(z01.RandIntBetween(1, 29), "qwertyuiopasdfghjklzxcvbnm")

		value := node{
			big:    big,
			little: little,
		}

		table = append(table, value)
	}

	for _, arg := range table {
		z01.Challenge(t, DoppelGanger, solutions.DoppelGanger, arg.big, arg.little)
	}
}
