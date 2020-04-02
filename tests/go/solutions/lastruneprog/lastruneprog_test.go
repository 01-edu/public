package main

import (
	"testing"

	"github.com/01-edu/z01"

	solutions "../../solutions"
)

func TestLastRuneProg(t *testing.T) {
	table := z01.MultRandASCII()
	table = append(table,
		"Hello!",
		"Salut!",
		"Ola!",
		z01.RandStr(z01.RandIntBetween(1, 15), z01.RandAlnum()),
	)
	for _, arg := range table {
		z01.Challenge(t, LastRune, solutions.LastRune, arg)
	}
}
