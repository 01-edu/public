package main

import (
	"testing"

	"github.com/01-edu/z01"

	solutions "../../solutions"
)

func TestFirstRuneProg(t *testing.T) {
	table := append(
		z01.MultRandASCII(),
		"Hello!",
		"Salut!",
		"Ola!",
		"♥01",
	)
	for _, arg := range table {
		z01.Challenge(t, FirstRune, solutions.FirstRune, arg)
	}
}
