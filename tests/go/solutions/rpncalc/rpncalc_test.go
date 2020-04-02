package main

import (
	"testing"

	"github.com/01-edu/z01"
)

func TestRpnCalc(t *testing.T) {
	args := []string{
		"1 2 * 3 * 4 +",
		"3 1 2 * * 4 %",
		"5 10 9 / - 50 *",
		"21 3 2 % 2 3 2 *",
		"1 2 3 4 +",
		"324   +    1 - 23 ",
		"32   / 22",
		"11 22 +",
		"23491234 102030932 -",
		"123 2222 /",
		"299   255 %",
		"15 76 *",
		"88 67 dks -",
		"     1      3 * 2 -",
	}

	for _, v := range args {
		z01.ChallengeMainExam(t, v)
	}
	z01.ChallengeMainExam(t)
	z01.ChallengeMainExam(t, "1 2 * 3 * 4 +", "10 33 - 12 %")
}
