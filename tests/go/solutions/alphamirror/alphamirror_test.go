package main

import (
	"testing"

	"github.com/01-edu/z01"
)

func TestAlphaMirror(t *testing.T) {
	arg1 := []string{""}
	arg2 := []string{"One", "ring!"}
	arg3 := []string{"testing spaces and #!*"}
	arg4 := []string{"more", "than", "three", "arguments"}
	arg5 := []string{"Upper anD LoWer cAsE"}
	arg6 := []string{z01.RandWords()}
	args := [][]string{arg1, arg2, arg3, arg4, arg5, arg6}

	for _, v := range args {
		z01.ChallengeMainExam(t, v...)
	}
}
