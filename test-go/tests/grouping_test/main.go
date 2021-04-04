package main

import (
	"github.com/01-edu/public/test-go/lib"
)

func validRegExp(n int) string {
	result := "("

	for i := 0; i < n; i++ {
		result += lib.RandStr(1, lib.Lower)
		if lib.RandInt()%2 == 0 {
			result += lib.RandStr(1, lib.Lower)
		}
		if i != n-1 {
			result += "|"
		}
	}

	result += ")"
	return result
}

func main() {
	args := [][2]string{
		{"(a)", "I'm heavyjumpsuit is on steady, Lighter when I'm lower, higher when I'm heavy"},
		{"(e|n)", "I currently have 4 windows opened up… and I don’t know why."},
		{"(hi)", "He swore he just saw his sushi move."},
		{"(s)", ""},
		{"i", "Something in the air"},
		{validRegExp(2), lib.RandStr(60, lib.Lower+lib.Space)},
		{validRegExp(2), lib.RandStr(60, lib.Lower+lib.Space)},
		{validRegExp(6), lib.RandStr(60, lib.Lower+lib.Space)},
		{lib.RandStr(1, "axyz"), lib.RandStr(10, "axyzdassbzzxxxyy cdq     ")},
	}
	for _, s := range args {
		lib.ChallengeMain("grouping", s[0], s[1])
	}
}
