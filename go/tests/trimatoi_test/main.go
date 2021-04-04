package main

import (
	student "student"

	"github.com/01-edu/public/go/tests/lib"
)

func trimAtoi(s string) int {
	runes := []rune(s)
	for i := range s {
		runes[i] = rune(s[i])
	}
	var numbers []rune
	for _, r := range runes {
		if (r >= '0' && r <= '9') || (len(numbers) == 0 && (r) == '-' || r == '+') {
			numbers = append(numbers, r)
		}
	}
	if len(numbers) == 0 || (len(numbers) == 1 && (numbers[0] == '-' || numbers[0] == '+')) {
		return 0
	}

	res, i, sign := 0, 0, 1

	if numbers[0] == '-' {
		sign = -1
		i++
	} else if numbers[0] == '+' {
		i++
	}
	for ; i < len(numbers); i++ {
		res = res*10 + int(numbers[i]) - '0'
	}

	return sign * res
}

func stringsToTrimAtoi(a []string) []string {
	alpha := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789012345678901234567890123456789"

	for index := 0; index < 4; index++ {
		s := ""
		s += lib.RandStr(lib.RandIntBetween(0, 2), alpha)
		x := lib.RandIntBetween(0, 14)
		if x <= 4 {
			s += "-"
		}
		s += lib.RandStr(lib.RandIntBetween(0, 10), alpha)
		a = append(a, s)
	}
	return a
}

func main() {
	a := []string{
		"",
		"12345",
		"str123ing45",
		"012 345",
		"Hello World!",
		"sd+x1fa2W3s4",
		"sd-x1fa2W3s4",
		"sdx1-fa2W3s4",
		lib.RandAlnum(),
	}
	a = stringsToTrimAtoi(a)
	for _, elem := range a {
		lib.Challenge("TrimAtoi", student.TrimAtoi, trimAtoi, elem)
	}
}

// TODO: refactor, including the subject
