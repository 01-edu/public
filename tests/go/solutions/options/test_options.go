package main

import (
	"strings"

	"github.com/01-edu/z01"
)

func main() {
	var table []string

	table = append(table, "-"+z01.RandLower(),
		" ",
		"-%",
		"-?",
		"-=",
		"-abc -ijk",
		"-z",
		"-abc -hijk",
		"-abcdefghijklmnopqrstuvwxyz",
		"-abcdefgijklmnopqrstuvwxyz",
		"-eeeeee",
		"-hhhhhh",
		"-h",
		"-hz",
		"-zh",
		"-z -h",
		strings.Join([]string{"-", z01.RandStr(10, z01.RuneRange('a', 'z'))}, ""),
	)

	for _, s := range table {
		z01.ChallengeMain("options", strings.Fields(s)...)
	}
}
