package main

import (
	"math/rand"
	"strings"

	"./student"
)

func split(s, sep string) []string {
	return strings.Split(s, sep)
}

func main() {
	separators := []string{"!=HA=!",
		"!==!",
		"    ",
		"|=choumi=|",
		"|<=>|",
		lib.RandStr(3, lib.RuneRange('A', 'Z')),
		"<<==123==>>",
		"[<>abc<>]"}

	type node struct {
		str string
		sep string
	}
	table := []node{}
	// 15 random slice of strings

	for i := 0; i < 15; i++ {
		separator := separators[rand.Intn(len(separators))]
		val := node{
			str: strings.Join(lib.MultRandAlnum(), separator),
			sep: separator,
		}
		table = append(table, val)
	}

	table = append(table,
		node{str: "HelloHAhowHAareHAyou?", sep: "HA"})

	for _, arg := range table {
		lib.Challenge("Split", student.Split, split, arg.str, arg.sep)
	}
}
