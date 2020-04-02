package main

import (
	"testing"

	"github.com/01-edu/z01"
)

func TestRot14Prog(t *testing.T) {
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
			z01.ChallengeMainExam(t, ""+s+"")
		}
	}
	z01.ChallengeMainExam(t, "", "something", "something1")
}
