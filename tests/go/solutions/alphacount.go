package solutions

import "unicode"

func AlphaCount(s string) (i int) {
	for _, r := range s {
		if unicode.IsLetter(r) {
			i++
		}
	}
	return i
}
