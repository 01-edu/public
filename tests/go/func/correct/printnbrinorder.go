package correct

import (
	"fmt"
	"sort"
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
		fmt.Print("0")
		return
	}
	digits := intToDigits(n)
	sort.Ints(digits)
	for _, i := range digits {
		fmt.Printf("%c", rune(i)+'0')
	}
	fmt.Println()
}
