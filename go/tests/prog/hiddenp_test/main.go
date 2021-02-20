package main

import (
	"github.com/01-edu/public/go/tests/lib"
)

func main() {
	args := [][2]string{
		{"fgex.;", "tyf34gdgf;'ektufjhgdgex.;.;rtjynur6"},
		{"abc", "2altrb53c.sse"},
		{"abc", "btarc"},
		{"DD", "DABC"},
		{""},
	}
	for i := 0; i < 30; i++ {
		args = append(args,
			[2]string{lib.RandStr(1, lib.Lower), lib.RandLower()},
			[2]string{lib.RandStr(1, lib.Upper), lib.RandUpper()},
		)
	}
	for _, v := range args {
		lib.ChallengeMain("hiddenp", v[0], v[1])
	}
}
