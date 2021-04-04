package main

import (
	"math/rand"
	"strings"

	student "student"

	"github.com/01-edu/public/test-go/lib"
)

func main() {
	separators := []string{
		"!=HA=!",
		"!==!",
		"    ",
		"|=choumi=|",
		"|<=>|",
		lib.RandStr(3, lib.RuneRange('A', 'Z')),
		"<<==123==>>",
		"[<>abc<>]",
	}

	type node struct {
		s   string
		sep string
	}
	table := []node{}
	// 15 random slice of strings

	for i := 0; i < 15; i++ {
		separator := separators[rand.Intn(len(separators))]
		val := node{
			s:   strings.Join(lib.MultRandAlnum(), separator),
			sep: separator,
		}
		table = append(table, val)
	}

	table = append(table,
		node{s: "HelloHAhowHAareHAyou?", sep: "HA"})

	for _, arg := range table {
		lib.Challenge("Split", student.Split, strings.Split, arg.s, arg.sep)
	}
}
