package correct

import "strconv"

func BasicAtoi2(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}
