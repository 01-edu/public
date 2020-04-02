package main

import (
	"strconv"
	"testing"

	"github.com/01-edu/z01"
)

type node struct {
	A, B, C, D, E, F int
}

func TestCostumeProfit(t *testing.T) {

	table := []node{}

	for i := 0; i < 25; i++ {
		a := z01.RandIntBetween(0, 1000)
		b := z01.RandIntBetween(0, 1000)
		c := z01.RandIntBetween(0, 1000)
		d := z01.RandIntBetween(0, 1000)
		e := z01.RandIntBetween(0, 1000)
		f := z01.RandIntBetween(0, 1000)
		table = append(table, node{a, b, c, d, e, f})
	}

	for _, arg := range table {
		a := strconv.Itoa(arg.A)
		b := strconv.Itoa(arg.B)
		c := strconv.Itoa(arg.C)
		d := strconv.Itoa(arg.D)
		e := strconv.Itoa(arg.E)
		f := strconv.Itoa(arg.F)
		z01.ChallengeMainExam(t, a, b, c, d, e, f)
	}
}
