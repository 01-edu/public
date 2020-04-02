package solutions

import (
	"github.com/01-edu/z01"
)

func sorting(arr []int) []int {
	for run := true; run; {
		run = false
		for i := 1; i < len(arr); i++ {
			if arr[i] < arr[i-1] {
				arr[i], arr[i-1] = arr[i-1], arr[i]
				run = true
			}
		}
	}
	return arr
}

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
	for _, c := range sorting(intToDigits(n)) {
		z01.PrintRune(rune(c) + '0')
	}
}
