package main

import (
	"github.com/01-edu/public/test-go/lib"
)

func main() {
	table := [][]string{{
		"? love",
		"? is",
		"A loveis",
		"? love",
		"? Who",
		"A Whoareyou",
		"? is",
	}}

	sets := [][]string{
		{"An", "array", "variable", "denotes", "the", "entire", "array"},
		{"This", "means", "that", "when", "you", "assign", "or", "pass"},
		{"To", "avoid", "the", "copy", "you", "could", "pass"},
		{"struct", "but", "with", "indexed", "rather", "than", "named", "fields"},
	}
	ops := []string{
		"?", "x", "A",
	}

	for i := 0; i < 6; i++ {
		result := []string{}
		nOps := lib.RandIntBetween(3, 15)
		index := lib.RandIntBetween(0, len(sets)-1)
		for j := 0; j < nOps; j++ {
			k := lib.RandIntBetween(0, len(ops)-1)
			s := lib.RandIntBetween(0, len(sets[index])-1)
			result = append(result, ops[k]+" "+sets[index][s])
		}
		table = append(table, result)
	}

	for _, arg := range table {
		lib.ChallengeMain("nenokku", arg...)
	}
}
