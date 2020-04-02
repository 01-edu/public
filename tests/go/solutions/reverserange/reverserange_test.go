package main

import (
	"strconv"
	"testing"

	"github.com/01-edu/z01"
)

func TestReverseRange(t *testing.T) {
	for i := 0; i < 10; i++ {
		start := z01.RandIntBetween(-20, 20)
		end := z01.RandIntBetween(-20, 20)
		z01.ChallengeMainExam(t, strconv.Itoa(start), strconv.Itoa(end))
	}
	z01.ChallengeMainExam(t, "2", "1", "3")
	z01.ChallengeMainExam(t, "a", "1")
	z01.ChallengeMainExam(t, "1", "b")
	z01.ChallengeMainExam(t, "1", "nan")
	z01.ChallengeMainExam(t, "nan", "b")
	z01.ChallengeMainExam(t)
}
