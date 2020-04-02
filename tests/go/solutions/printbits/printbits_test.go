package main

import (
	"strconv"
	"strings"
	"testing"

	"github.com/01-edu/z01"
)

func TestPrintBits(t *testing.T) {
	var arg []string
	for i := 0; i < 20; i++ {
		arg = append(arg, strconv.Itoa(z01.RandIntBetween(0, 255)))
	}
	arg = append(arg, "")
	arg = append(arg, "a")
	arg = append(arg, "bc")
	arg = append(arg, "def")
	arg = append(arg, "notanumber")
	arg = append(arg, z01.RandBasic())

	for _, v := range arg {
		z01.ChallengeMainExam(t, strings.Fields(v)...)
	}
}
