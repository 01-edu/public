package main

import (
	"strings"

	"github.com/01-edu/public/test-go/lib"
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
