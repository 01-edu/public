package main

import "github.com/01-edu/z01"

func validRegExp(n int) string {
	result := "("

	for i := 0; i < n; i++ {
		result += z01.RandStr(1, z01.Lower)
		if z01.RandInt()%2 == 0 {
			result += z01.RandStr(1, z01.Lower)
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
		{validRegExp(2), z01.RandStr(60, z01.Lower+z01.Space)},
		{validRegExp(2), z01.RandStr(60, z01.Lower+z01.Space)},
		{validRegExp(6), z01.RandStr(60, z01.Lower+z01.Space)},
		{z01.RandStr(1, "axyz"), z01.RandStr(10, "axyzdassbzzxxxyy cdq     ")},
	}
	for _, s := range args {
		z01.ChallengeMain("grouping", s[0], s[1])
	}
}
