package main

import (
	"strings"

	"github.com/01-edu/z01"
)

func main() {
	args := append(z01.MultRandWords(),
		"padinton paqefwtdjetyiytjneytjoeyjnejeyj",
		"ddf6vewg64f  twthgdwthdwfteewhrtag6h4ffdhsd",
		"abcdefghij efghijlmnopq",
		"123456 456789",
		"1 1",
		"1 2",
	)

	for i := 0; i < 5; i++ {
		s1 := z01.RandAlnum()
		s2 := z01.RandAlnum() + s1 + z01.RandAlnum()
		args = append(args,
			s1+" "+s2,
			z01.RandAlnum()+" "+z01.RandAlnum(),
		)
	}

	for _, s := range args {
		z01.ChallengeMain("inter", strings.Fields(s)...)
	}
}
