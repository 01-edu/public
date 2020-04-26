package solutions

import "strings"

func DoppelGanger(s, substr string) int {
	return strings.LastIndex(s, substr)
}
