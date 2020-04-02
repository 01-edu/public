package main

import (
	"testing"

	"github.com/01-edu/z01"
)

func TestLastWord(t *testing.T) {
	args := []string{
		"FOR PONY",
		"this ... is sparta, then again, maybe not",
		" ",
		" lorem,ipsum ",
	}

	for i := 0; i < 5; i++ {
		args = append(args, z01.RandWords())
	}

	for _, v := range args {
		z01.ChallengeMainExam(t, v)
	}

	z01.ChallengeMainExam(t, "a", "b")
	z01.ChallengeMainExam(t)
}
