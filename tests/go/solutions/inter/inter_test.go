package main

import (
	"strings"
	"testing"

	"github.com/01-edu/z01"
)

func TestInter(t *testing.T) {
	table := append(z01.MultRandWords(),
		"padinton paqefwtdjetyiytjneytjoeyjnejeyj",
		"ddf6vewg64f  twthgdwthdwfteewhrtag6h4ffdhsd",
	)
	table = append(table, "abcdefghij efghijlmnopq")
	table = append(table, "123456 456789")
	table = append(table, "1 1")
	table = append(table, "1 2")

	for i := 0; i < 5; i++ {
		str1 := z01.RandAlnum()
		str2 := strings.Join([]string{z01.RandAlnum(), str1, z01.RandAlnum()}, "")

		table = append(table, strings.Join([]string{str1, str2}, " "))
		table = append(table, strings.Join([]string{z01.RandAlnum(), z01.RandAlnum()}, " "))

	}

	for _, s := range table {
		z01.ChallengeMainExam(t, strings.Fields(s)...)
	}
}
