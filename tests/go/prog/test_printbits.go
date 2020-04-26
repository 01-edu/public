package main

import (
	"strconv"
	"strings"

	"../lib"
)

func main() {
	var arg []string
	for i := 0; i < 20; i++ {
		arg = append(arg, strconv.Itoa(lib.RandIntBetween(0, 255)))
	}
	arg = append(arg, "")
	arg = append(arg, "a")
	arg = append(arg, "bc")
	arg = append(arg, "def")
	arg = append(arg, "notanumber")
	arg = append(arg, lib.RandBasic())

	for _, v := range arg {
		lib.ChallengeMain("printbits", strings.Fields(v)...)
	}
}
