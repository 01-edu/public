package main

import (
	"strings"
)

func ContainsPlusMinus(s string) bool {
	for _, c := range s {
		if c == '+' || c == '-' {
			return true
		}
	}
	return false
}

func UniqueChar(s string) bool {
	r := []rune(s)
	n := len(r)
	for i := 0; i < n-1; i++ {
		if InStr(r[i], s[i+1:n]) {
			return false
		}

	}
	return true
}

func InStr(c rune, s string) bool {
	for _, r := range s {
		if c == r {
			return true
		}
	}
	return false
}

func ValidBase(base string) bool {
	return len(base) >= 2 && !ContainsPlusMinus(base) && UniqueChar(base)
}

func Power(nbr int, pwr int) int {
	if pwr == 0 {
		return 1
	}
	if pwr == 1 {
		return nbr
	}
	return nbr * Power(nbr, pwr-1)

}

func Index(s string, toFind string) int {
	result := strings.Index(s, toFind)
	return result
}

func AtoiBase(s string, base string) int {
	var result int
	var i int
	sign := 1
	lengthBase := len(base)
	lengths := len(s)

	if !ValidBase(base) {
		return 0
	}
	if s[i] == '-' {
		sign = -1
	}
	if s[i] == '-' || s[i] == '+' {
		i++
	}
	for i < len(s) {
		result = result + (Index(base, string(s[i])) * Power(lengthBase, lengths-1))
		i++
		lengths--
	}
	return result * sign
}

func main() {

}
