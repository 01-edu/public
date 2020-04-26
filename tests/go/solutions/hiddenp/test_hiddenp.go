package main

import "github.com/01-edu/z01"

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
			[2]string{z01.RandStr(1, z01.Lower), z01.RandLower()},
			[2]string{z01.RandStr(1, z01.Upper), z01.RandUpper()},
		)
	}
	for _, v := range args {
		z01.ChallengeMain("hiddenp", v[0], v[1])
	}
}
