package correct

import "strconv"

func BasicAtoi(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}
