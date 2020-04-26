package solutions

import (
	"sort"

	"github.com/01-edu/z01"
)

func intToDigits(n int) (digits []int) {
	for n > 0 {
		if n == 0 {
			digits = append(digits, 0)
		} else {
			digits = append(digits, n%10)
		}
		n /= 10
	}
	return
}

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	digits := intToDigits(n)
	sort.Ints(digits)
	for _, i := range digits {
		z01.PrintRune(rune(i) + '0')
	}
}
