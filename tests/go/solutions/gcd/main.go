package main

import (
	"fmt"
	"os"
	"strconv"
)

//Greatest common divisor
func gcd(num1, num2 uint) uint {

	for i := num1; i > 0; i-- {
		if num1%i == 0 && num2%i == 0 {
			return i
		}
	}

	return 1
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println()
		os.Exit(0)
	}

	v1, _ := strconv.Atoi(os.Args[1])

	v2, _ := strconv.Atoi(os.Args[2])

	fmt.Println(gcd(uint(v1), uint(v2)))

}
