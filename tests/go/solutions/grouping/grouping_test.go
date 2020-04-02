package main

import (
	"testing"

	"github.com/01-edu/z01"
)

func TestGrouping(t *testing.T) {

	type args struct {
		first  string
		second string
	}

	arr := []args{
		{first: "(a)",
			second: "I'm heavyjumpsuit is on steady, Lighter when I'm lower, higher when I'm heavy"},

		{first: "(e|n)", second: "I currently have 4 windows opened up… and I don’t know why."},
		{first: "(hi)", second: "He swore he just saw his sushi move."},
		{first: "(s)", second: ""},
		{first: "i", second: "Something in the air"},
	}

	for i := 0; i < 2; i++ {
		helper := args{first: validRegExp(2), second: validString(60)}
		arr = append(arr, helper)
	}
	helper := args{first: validRegExp(6), second: validString(60)}
	arr = append(arr, helper)

	helper = args{first: z01.RandStr(1, "axyz"), second: z01.RandStr(10, "axyzdassbzzxxxyy cdq     ")}
	arr = append(arr, helper)

	for _, s := range arr {
		z01.ChallengeMainExam(t, s.first, s.second)
	}
}

func validString(len int) string {
	s := z01.RandStr(len, "abcdefijklmnopqrstyz           ")
	return s
}

func validRegExp(n int) string {
	result := "("

	for i := 0; i < n; i++ {
		result += string(z01.RandStr(1, "abcdefijklmnopqrstyz"))
		if z01.RandInt()%2 == 0 {
			result += string(z01.RandStr(1, "abcdefijklmnopqrstyz"))
		}
		if i != n-1 {
			result += "|"
		}
	}

	result += ")"
	return result
}
