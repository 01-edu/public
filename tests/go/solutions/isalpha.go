package solutions

import "unicode"

func IsAlpha(s string) bool {
	for _, r := range s {
		if !unicode.Is(unicode.Hex_Digit, r) {
			return false
		}
	}
	return true
}
