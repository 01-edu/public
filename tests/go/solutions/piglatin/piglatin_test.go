package main

import (
	"testing"

	"github.com/01-edu/z01"
)

func TestPigLatin(t *testing.T) {
	type args struct {
		first []string
	}
	arr := []args{{first: []string{"", "pig", "is", "crunch", "crnch"}}, {first: []string{"something", "else"}}}

	for i := 0; i < 4; i++ {
		arr = append(arr, args{first: z01.MultRandBasic()})
	}

	for _, v := range arr {
		z01.ChallengeMainExam(t, v.first...)
	}
}
