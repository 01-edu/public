package correct

import "strings"

func DoppelGanger(s, substr string) int {
	return strings.LastIndex(s, substr)
}
