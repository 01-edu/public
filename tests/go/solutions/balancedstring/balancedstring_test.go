package main

import (
	"testing"

	"github.com/01-edu/z01"
)

func TestBalancedString(t *testing.T) {

	// Declaration of an empty array of type string
	table := []string{}

	// Filing of this array with the tests array provided.
	table = append(table,
		"CDCCDDCDCD",
		"CDDDDCCCDC",
		"DDDDCCCC",
		"CDCCCDDCDD",
		"CDCDCDCDCDCDCDCD",
		"CCCDDDCDCCCCDC",
		"DDDDDDDDDDDDDDDDDDDDDDDDCCCCCCCCCCCCCCCCCCCCCCCC",
		"DDDDCDDDDCDDDCCC",
	)

	for i := 0; i < 15; i++ {
		value := ""
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
					value += letter
				} else if tmpD > 0 && letter == "D" {
					tmpD--
					value += letter
				}
			}

			tmpC = countC
			tmpD = countD
			for tmpC > 0 || tmpD > 0 {
				letter := z01.RandStr(1, "CD")
				if tmpC > 0 && letter == "D" {
					tmpC--
					value += letter
				} else if tmpD > 0 && letter == "C" {
					tmpD--
					value += letter
				}
			}
		}
		table = append(table, value)
	}

	//The table with 8 specific exercises and 15 randoms is now ready to be "challenged"
	//This time, contrary to the nauuo exercise,
	//We can use the ChallengeMainExam function.

	for _, arg := range table {
		z01.ChallengeMainExam(t, arg)
	}
}
