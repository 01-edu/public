package main

import (
	"github.com/01-edu/z01"
)

func main() {
	type nodeTest struct {
		data []string
	}

	table := []nodeTest{}
	for i := 0; i < 2; i++ {
		val := nodeTest{
			data: z01.MultRandWords(),
		}
		table = append(table, val)
	}

	for _, arg := range table {
		for _, s := range arg.data {
			z01.ChallengeMain("" + s + "")
		}
	}
	z01.ChallengeMain("", "something", "something1")
}
