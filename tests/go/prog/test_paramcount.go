package main

import (
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	args := [][]string{
		{"2", "5", "u", "19"},
		{"2"},
		{"1", "2", "3", "5", "7", "24"},
		{"6", "12", "24"},
	}

	for i := 0; i < 10; i++ {
		var arg []string
		init := z01.RandIntBetween(0, 10)
		for i := init; i < init+z01.RandIntBetween(5, 10); i++ {
			arg = append(arg, strconv.Itoa(i))
		}
		args = append(args, arg)
	}

	for i := 0; i < 1; i++ {
		args = append(args, z01.MultRandWords())
	}

	for _, v := range args {
		z01.ChallengeMain("paramcount", v...)
	}

	z01.ChallengeMain("paramcount")
}
