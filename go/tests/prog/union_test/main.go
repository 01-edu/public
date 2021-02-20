package main

import (
	"strings"

	"github.com/01-edu/public/go/tests/lib"
)

func main() {
	s1 := lib.RandAlnum()
	s2 := strings.Join([]string{lib.RandAlnum(), s1, lib.RandAlnum()}, "")

	args := [][]string{
		{"zpadinton", "paqefwtdjetyiytjneytjoeyjnejeyj"},
		{"ddf6vewg64f", "gtwthgdwthdwfteewhrtag6h4ffdhsd"},
		{""},
		{"rien", "cette phrase ne cache rien"},
		{" this is ", " wait shr"},
		{" more ", "then", "two", "arguments"},
		{s1, s2},
		{lib.RandAlnum(), lib.RandAlnum()},
	}

	for _, v := range args {
		lib.ChallengeMain("union", v...)
	}
}
