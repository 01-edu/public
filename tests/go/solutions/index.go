package solutions

import (
	"strings"
)

func Index(s string, toFind string) int {
	result := strings.Index(s, toFind)
	return result
}
