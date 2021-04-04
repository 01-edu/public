package main

import (
	student "student"

	"github.com/01-edu/public/test-go/lib"
	"github.com/01-edu/public/test-go/lib/base"
)

func main() {
	type node struct {
		s    string
		base string
	}

	table := []node{}

	// 15 random pairs of string numbers with valid bases
	for i := 0; i < 15; i++ {
		validBaseToInput := base.Valid()
		val := node{
			s:    base.StringFrom(validBaseToInput),
			base: validBaseToInput,
		}
		table = append(table, val)
	}
	// 15 random pairs of string numbers with invalid bases
	for i := 0; i < 15; i++ {
		invalidBaseToInput := base.Invalid()
		val := node{
			s:    "thisinputshouldnotmatter",
			base: invalidBaseToInput,
		}
		table = append(table, val)
	}
	table = append(table,
		node{s: "125", base: "0123456789"},
		node{s: "1111101", base: "01"},
		node{s: "7D", base: "0123456789ABCDEF"},
		node{s: "uoi", base: "choumi"},
		node{s: "bbbbbab", base: "-ab"},
	)
	for _, arg := range table {
		lib.Challenge("AtoiBase", student.AtoiBase, base.Atoi, arg.s, arg.base)
	}
}
