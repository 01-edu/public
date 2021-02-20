package main

import (
	"strconv"

	"github.com/01-edu/public/go/tests/lib"
)

func main() {
	args := [][]string{
		{"1"},
		{"2"},
		{"3"},
		{"4"},
		{"1024"},
		{"4096"},
		{"8388608"},
		{"1", "2"},
		{},
	}
	for i := 0; i < 12; i++ {
		args = append(args, []string{strconv.Itoa(lib.RandIntBetween(1, 2048))})
	}
	for _, v := range args {
		lib.ChallengeMain("ispowerof2", v...)
	}
}
