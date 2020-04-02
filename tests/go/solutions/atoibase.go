package solutions

import (
	"math/rand"

	"github.com/01-edu/z01"
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

// this function is used to create the tests (input VALID bases here)

func RandomValidBase() string {
	validBases := []string{
		"01",
		"CHOUMIisDAcat!",
		"choumi",
		"0123456789",
		"abc", "Zone01",
		"0123456789ABCDEF",
		"WhoAmI?",
	}
	index := rand.Intn(len(validBases))
	return validBases[index]
}

// this function is used to create the tests (input INVALID bases here)
func RandomInvalidBase() string {
	invalidBases := []string{
		"0",
		"1",
		"CHOUMIisdacat!",
		"choumiChoumi",
		"01234567890",
		"abca",
		"Zone01Zone01",
		"0123456789ABCDEF0",
		"WhoAmI?IamWhoIam",
	}
	index := z01.RandIntBetween(0, len(invalidBases)-1)
	return invalidBases[index]
}

// this function is used to create the random STRING number from VALID BASES
func RandomStringFromBase(base string) string {
	letters := []rune(base)
	size := z01.RandIntBetween(1, 10)
	r := make([]rune, size)
	for i := range r {
		r[i] = letters[rand.Intn(len(letters))]
	}
	return string(r)
}
