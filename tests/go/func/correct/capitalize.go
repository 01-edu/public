package correct

import (
	"strings"
	"unicode"
)

func Capitalize(s string) string {
	r := []rune(strings.ToLower(s))

	if unicode.IsLower(r[0]) {
		r[0] = unicode.ToUpper(r[0])
	}

	for i := 1; i < len(r); i++ {
		if !unicode.Is(unicode.ASCII_Hex_Digit, r[i-1]) && unicode.IsLower(r[i]) {
			r[i] = unicode.ToUpper(r[i])
		}
	}
	return string(r)
}
