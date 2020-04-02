package main

import (
	"strconv"
	"testing"

	"github.com/01-edu/z01"
)

func TestGCD(t *testing.T) {
	arg1 := []string{"23"}
	arg2 := []string{"12", "23"}
	arg3 := []string{"25", "15"}
	arg4 := []string{"23043", "122"}
	arg5 := []string{"11", "77"}
	args := [][]string{arg1, arg2, arg3, arg4, arg5}

	for i := 0; i < 25; i++ {

		number1 := strconv.Itoa(z01.RandIntBetween(1, 100000))
		number2 := strconv.Itoa(z01.RandIntBetween(1, 100))
		args = append(args, []string{number1, number2})
	}

	for _, v := range args {
		z01.ChallengeMainExam(t, v...)
	}
}
