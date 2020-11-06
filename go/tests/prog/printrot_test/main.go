package main

import (
	"lib"
	"strings"
)

func main() {
	table := string{}
	table = append(table, "A")
	table = append(table, "a")
	table = append(table, "b")
	table = append(table, "c")
	table = append(table, "d")
	table = append(table, "e")
	table = append(table, "f")
	table = append(table, "g")
	table = append(table, "h")
	table = append(table, "i")
	table = append(table, "j")
	table = append(table, "k")
	table = append(table, "l")
	table = append(table, "m")
	table = append(table, "n")
	table = append(table, "o")
	table = append(table, "p")
	table = append(table, "q")
	table = append(table, "r")
	table = append(table, "s")
	table = append(table, "t")
	table = append(table, "u")
	table = append(table, "v")
	table = append(table, "w")
	table = append(table, "x")
	table = append(table, "y")
	table = append(table, "z")
	table = append(table, "Z")
	table = append(table, "2 arguments")
	table = append(table, "4 arguments so invalid")

	for _, s := range table {
		lib.ChallengeMain("printrot", strings.Fields(s)...)
	}
}
