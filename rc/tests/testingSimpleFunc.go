package solutions

import (
	"regexp"
)

func SimpleFunc(str string) int {
	re := regexp.MustCompile(`[a-zA-Z]`)
	found := re.FindAll([]byte(str), -1)
	return len(found)
}
