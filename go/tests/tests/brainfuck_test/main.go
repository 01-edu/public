package main

import (
	"strings"

	"github.com/01-edu/public/go/tests/lib"
)

func main() {
	// individual tests 1)Hello World! 2)Hi 3)abc 4)ABC

	args := []string{
		"++++++++++[>+++++++>++++++++++>+++>+<<<<-]>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.--------.>+.>.",
		"+++++[>++++[>++++H>+++++i<<-]>>>++\n<<<<-]>>--------.>+++++.>.",
		"++++++++++[>++++++++++>++++++++++>++++++++++<<<-]>---.>--.>-.>++++++++++.",
		"ld++++++++++++++++++++++++++++++++++++++++++++this+is++a++++comment++++++++++++++[>d+<-]>.+.+.>++++++++++.",
		strings.Join([]string{"ld++++++++++++++++++++++++++++++++++++++++++++this+is++a++++comment++++++++++++++[>d+<-]>.+", lib.RandStr(lib.RandIntBetween(1, 10), ".+"), ".+.>++++++++++."}, ""),
	}
	for _, v := range args {
		lib.ChallengeMain("brainfuck", v)
	}
}
