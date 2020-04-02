package solutions

import (
	"github.com/01-edu/z01"
)

func PrintNbrBase(n int, base string) {
	if ValidBase(base) {
		length := len(base)
		sign := 1
		rbase := []rune(base)
		if n < 0 {
			z01.PrintRune('-')
			sign = -1
		}
		if n < length && n >= 0 {
			z01.PrintRune(rbase[n])
		} else {
			PrintNbrBase(sign*(n/length), base)
			z01.PrintRune(rbase[sign*(n%length)])
		}

	} else {
		z01.PrintRune('N')
		z01.PrintRune('V')
	}
}
