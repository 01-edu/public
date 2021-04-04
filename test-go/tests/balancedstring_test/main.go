package main

import (
	"github.com/01-edu/public/test-go/lib"
)

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
		chunks := lib.RandIntBetween(5, 10)
		for j := 0; j < chunks; j++ {
			countC := lib.RandIntBetween(1, 5)
			countD := lib.RandIntBetween(1, 5)
			tmpC := countC
			tmpD := countD
			for tmpC > 0 || tmpD > 0 {
				letter := lib.RandStr(1, "CD")
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
				letter := lib.RandStr(1, "CD")
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
		lib.ChallengeMain("balancedstring", arg)
	}
}
