package main

import (
	"github.com/01-edu/z01"
)

func main() {
	table := z01.MultRandASCII()
	for _, arg := range table {
		z01.ChallengeMain("printstrprog", arg)
	}
	z01.ChallengeMain("printstrprog", "Hello World!")
}
