package main

import (
	"strings"

	"github.com/01-edu/public/go/tests/lib"
)

func main() {
	var table []string

	table = append(table, "-"+lib.RandLower(),
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
		strings.Join([]string{"-", lib.RandStr(10, lib.RuneRange('a', 'z'))}, ""),
	)

	for _, s := range table {
		lib.ChallengeMain("options", strings.Fields(s)...)
	}
}
