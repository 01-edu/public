package main

import (
	"fmt"
	"os"
	"strconv"
)

// Greatest common divisor
func gcd(num1, num2 int) int {
	for i := num1; i > 0; i-- {
		if num1%i == 0 && num2%i == 0 {
			return i
		}
	}
	return 1
}

func main() {
	if len(os.Args) != 3 {
		return
	}
	v1, _ := strconv.Atoi(os.Args[1])
	v2, _ := strconv.Atoi(os.Args[2])
	fmt.Println(gcd(v1, v2))
}
