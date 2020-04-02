package solutions

import (
	"regexp"
)

func AlphaCount(str string) int {
	re := regexp.MustCompile(`[a-zA-Z]`)
	found := re.FindAll([]byte(str), -1)
	return len(found)
}
