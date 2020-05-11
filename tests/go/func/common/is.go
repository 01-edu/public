package common

import (
	"math/big"
	"unicode"
)

func IsNumeric(s string) bool {
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

func IsLower(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) {
			return false
		}
	}
	return true
}

func IsUpper(s string) bool {
	for _, r := range s {
		if !(r >= 'A' && r <= 'Z') {
			return false
		}
	}
	return true
}

func IsPrime(i int) bool {
	return big.NewInt(int64(i)).ProbablyPrime(0)
}
