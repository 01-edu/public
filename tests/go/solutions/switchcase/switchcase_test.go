package main

import (
	"testing"

	"github.com/01-edu/z01"
)

func TestSwitchcase(t *testing.T) {
	table := append(z01.MultRandWords(), " ")
	table = append(table, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	table = append(table, "abcdefghi jklmnop qrstuvwxyz ABCDEFGHI JKLMNOPQR STUVWXYZ ! ")

	for _, s := range table {
		z01.ChallengeMainExam(t, s)
	}
}
