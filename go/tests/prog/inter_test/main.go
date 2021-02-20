package main

import (
	"strings"

	"github.com/01-edu/public/go/tests/lib"
)

func main() {
	args := append(lib.MultRandWords(),
		"padinton paqefwtdjetyiytjneytjoeyjnejeyj",
		"ddf6vewg64f  twthgdwthdwfteewhrtag6h4ffdhsd",
		"abcdefghij efghijlmnopq",
		"123456 456789",
		"1 1",
		"1 2",
	)

	for i := 0; i < 5; i++ {
		s1 := lib.RandAlnum()
		s2 := lib.RandAlnum() + s1 + lib.RandAlnum()
		args = append(args,
			s1+" "+s2,
			lib.RandAlnum()+" "+lib.RandAlnum(),
		)
	}

	for _, s := range args {
		lib.ChallengeMain("inter", strings.Fields(s)...)
	}
}
