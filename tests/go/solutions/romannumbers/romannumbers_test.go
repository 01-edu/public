package main

import (
	"strconv"
	"testing"

	"github.com/01-edu/z01"
)

func TestRomanNumbers(t *testing.T) {
	rand := []string{
		"0",
		"4000",
		"5000",
		"12433",
		"hello",
		"good luck",
	}
	for i := 0; i < 7; i++ {
		rand = append(rand, strconv.Itoa(z01.RandIntBetween(0, 4000)))
	}
	for _, v := range rand {
		z01.ChallengeMainExam(t, v)
	}
}
