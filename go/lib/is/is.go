package is

import (
	"math/big"
	"unicode"
)

func Digit(s string) bool {
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

func Lower(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) {
			return false
		}
	}
	return true
}

func Upper(s string) bool {
	for _, r := range s {
		if !(r >= 'A' && r <= 'Z') {
			return false
		}
	}
	return true
}

func Prime(i int) bool {
	return big.NewInt(int64(i)).ProbablyPrime(0)
}
