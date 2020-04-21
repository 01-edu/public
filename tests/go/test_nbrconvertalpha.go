package main

import (
	"strconv"
	"strings"

	"github.com/01-edu/z01"
)

func main() {
	type node struct {
		array []string
	}

	table := []node{}
	for i := 0; i < 5; i++ {
		m := z01.MultRandIntBetween(1, 46)
		s := ""
		str := []string{}
		for _, j := range m {
			s += strconv.Itoa(j) + " "
		}
		str = append(str, s)
		v := node{
			array: str,
		}
		table = append(table, v)
	}

	table = append(table, node{
		array: []string{
			"--upper 8 5 25",
			"8 5 12 12 15",
			"12 5 7 5 14 56 4 1 18 25",
			"12 5 7 5 14 56 4 1 18 25 32 86 h",
			"32 86 h",
			""},
	})

	for _, i := range table {
		for _, a := range i.array {
			z01.ChallengeMain(strings.Fields((a))...)
		}
	}
}
