package solutions

import (
	"strings"
)

func DoppelGanger(big, little string) int {
	return strings.LastIndex(big, little)
}
