package main

import "github.com/01-edu/z01"

func main() {
	table := []string{
		"CDCCDDCDCD",
		"CDDDDCCCDC",
		"DDDDCCCC",
		"CDCCCDDCDD",
		"CDCDCDCDCDCDCDCD",
		"CCCDDDCDCCCCDC",
		"DDDDDDDDDDDDDDDDDDDDDDDDCCCCCCCCCCCCCCCCCCCCCCCC",
		"DDDDCDDDDCDDDCCC",
	}

	for i := 0; i < 15; i++ {
		s := ""
		chunks := z01.RandIntBetween(5, 10)
		for j := 0; j < chunks; j++ {
			countC := z01.RandIntBetween(1, 5)
			countD := z01.RandIntBetween(1, 5)
			tmpC := countC
			tmpD := countD
			for tmpC > 0 || tmpD > 0 {
				letter := z01.RandStr(1, "CD")
				if tmpC > 0 && letter == "C" {
					tmpC--
					s += letter
				} else if tmpD > 0 && letter == "D" {
					tmpD--
					s += letter
				}
			}

			tmpC = countC
			tmpD = countD
			for tmpC > 0 || tmpD > 0 {
				letter := z01.RandStr(1, "CD")
				if tmpC > 0 && letter == "D" {
					tmpC--
					s += letter
				} else if tmpD > 0 && letter == "C" {
					tmpD--
					s += letter
				}
			}
		}
		table = append(table, s)
	}

	for _, arg := range table {
		z01.ChallengeMain("balancedstring", arg)
	}
}
