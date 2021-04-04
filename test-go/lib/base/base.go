package base

import (
	"math/rand"
	"strings"

	"github.com/01-edu/public/test-go/lib"
)

func uniqueChar(s string) bool {
	m := map[rune]struct{}{}
	for _, r := range s {
		if _, ok := m[r]; ok {
			return false
		}
		m[r] = struct{}{}
	}
	return true
}

func IsValid(base string) bool {
	return len(base) >= 2 && !strings.ContainsAny(base, "+-") && uniqueChar(base)
}

func power(nbr int, pwr int) int {
	if pwr == 0 {
		return 1
	}
	if pwr == 1 {
		return nbr
	}
	return nbr * power(nbr, pwr-1)
}

func Atoi(s string, base string) int {
	var result int
	var i int
	sign := 1
	lengthBase := len(base)
	lengths := len(s)

	if !IsValid(base) {
		return 0
	}
	if s[i] == '-' {
		sign = -1
	}
	if s[i] == '-' || s[i] == '+' {
		i++
	}
	for i < len(s) {
		result += strings.Index(base, string(s[i])) * power(lengthBase, lengths-1)
		i++
		lengths--
	}
	return result * sign
}

func Valid() string {
	valid := []string{
		"01",
		"CHOUMIisDAcat!",
		"choumi",
		"0123456789",
		"abc",
		"Zone01",
		"0123456789ABCDEF",
		"WhoAmI?",
	}
	i := rand.Intn(len(valid))
	return valid[i]
}

func Invalid() string {
	invalid := []string{
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
	i := rand.Intn(len(invalid))
	return invalid[i]
}

func StringFrom(base string) string {
	letters := []rune(base)
	size := lib.RandIntBetween(1, 10)
	runes := make([]rune, size)
	for i := range runes {
		runes[i] = letters[rand.Intn(len(letters))]
	}
	return string(runes)
}

func ConvertNbr(n int, base string) string {
	var result string
	length := len(base)

	for n >= length {
		result = string(base[(n%length)]) + result
		n = n / length
	}
	result = string(base[n]) + result

	return result
}

func Convert(nbr, baseFrom, baseTo string) string {
	resultIntermediary := Atoi(nbr, baseFrom)

	resultFinal := ConvertNbr(resultIntermediary, baseTo)

	return resultFinal
}
