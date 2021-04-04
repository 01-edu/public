package main

import (
	"strconv"
	"strings"

	"github.com/01-edu/public/test-go/lib"
)

func main() {
	var args []string
	for i := 0; i < 20; i++ {
		args = append(args, strconv.Itoa(lib.RandIntBetween(0, 255)))
	}
	args = append(args,
		"",
		"a",
		"bc",
		"def",
		"notanumber",
		lib.RandBasic(),
	)
	for _, v := range args {
		lib.ChallengeMain("printbits", strings.Fields(v)...)
	}
}
