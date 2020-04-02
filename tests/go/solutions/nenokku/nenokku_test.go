package main

import (
	"testing"
	"github.com/01-edu/z01"
)

type node struct {
	operations []string
}

func TestNenokku(t *testing.T) {

	table := []node{}

	table = append(table,
		node{
			[]string{
				"? love",
				"? is",
				"A loveis",
				"? love",
				"? Who",
				"A Whoareyou",
				"? is",
			},
		},
	)

	sets := [][]string {
		[]string{"An", "array", "variable", "denotes", "the", "entire", "array"},
		[]string{"This", "means", "that", "when", "you", "assign", "or", "pass"},
		[]string{"To", "avoid", "the", "copy", "you", "could", "pass"},
		[]string{"struct", "but", "with", "indexed", "rather", "than", "named", "fields"},
	}
	ops := []string {
		"?", "x", "A",
	}

	for i := 0; i < 6; i++ {
		result := []string{}
		nOps := z01.RandIntBetween(3, 15)
		index := z01.RandIntBetween(0, len(sets)-1)
		for j := 0; j < nOps; j++ {
			k := z01.RandIntBetween(0, len(ops)-1)
			s := z01.RandIntBetween(0, len(sets[index])-1)
			result = append(result, ops[k] + " " + sets[index][s])
		}
		table = append(table, node{result})
	}


	for _, arg := range table {	
		z01.ChallengeMainExam(t, arg.operations...)
	}
}
