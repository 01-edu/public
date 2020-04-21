package main

import (
	"github.com/01-edu/z01"
)

func main() {
	type node struct {
		s         string
		toCompare string
	}

	table := []node{}

	// the first 7 values are returning 0 for this test
	for i := 0; i < 7; i++ {
		wordToTest := z01.RandASCII()

		val := node{
			s:         wordToTest,
			toCompare: wordToTest,
		}
		table = append(table, val)
	}

	// the next 7 values are supposed to return 1 or -1 for this test
	for i := 0; i < 7; i++ {
		wordToTest := z01.RandASCII()
		wrongMatch := z01.RandASCII()

		val := node{
			s:         wordToTest,
			toCompare: wrongMatch,
		}
		table = append(table, val)
	}
	// those are the test values from the README examples
	table = append(table,
		node{s: "Hello!", toCompare: "Hello!"},
		node{s: "Salut!", toCompare: "lut!"},
		node{s: "Ola!", toCompare: "Ol"},
	)
	for _, arg := range table {
		z01.ChallengeMain("compareprog", arg.s, arg.toCompare)
	}
	z01.ChallengeMain("compareprog")
	z01.ChallengeMain("compareprog", "1 arg", "2args", "3args")
}
