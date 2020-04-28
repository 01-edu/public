package main

import (
	"strings"

	"../../lib"
)

func main() {
	table := append(lib.MultRandWords(),
		" ",
		"faya fgvvfdxcacpolhyghbreda",
		"faya fgvvfdxcacpolhyghbred",
		"error rrerrrfiiljdfxjyuifrrvcoojh",
		"error rrerrrfiiljdfxjyuifrrvcoojrh",
	)

	for _, s := range table {
		lib.ChallengeMain("wdmatch", strings.Fields(s)...)
	}
}
