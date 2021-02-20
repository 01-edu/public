package main

import (
	"strconv"
	"strings"

	"github.com/01-edu/public/go/tests/lib"
)

func main() {
	operatorsTable := []string{"+", "-", "*", "/", "%"}

	table := []string{}

	for i := 0; i < 4; i++ {
		firstArg := strconv.Itoa(lib.RandIntBetween(-1000, 1000))
		secondArg := strconv.Itoa(lib.RandIntBetween(0, 1000))

		for _, operator := range operatorsTable {
			table = append(table, firstArg+" "+operator+" "+secondArg)

		}
	}

	table = append(table, "1 + 1")
	table = append(table, "hello + 1")
	table = append(table, "1 p 1")
	table = append(table, "1 / 0")
	table = append(table, "1 # 1")
	table = append(table, "1 % 0")
	table = append(table, "1 * 1")
	table = append(table, "1argument")
	table = append(table, "2 arguments")
	table = append(table, "4 arguments so invalid")
	table = append(table, "9223372036854775807 + 1")
	table = append(table, "9223372036854775809 - 3")
	table = append(table, "9223372036854775807 * 3")
	for _, s := range table {
		lib.ChallengeMain("doop", strings.Fields(s)...)
	}
}
