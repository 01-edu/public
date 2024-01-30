package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.DigitLen(100, 10))
	fmt.Println(piscine.DigitLen(100, 2))
	fmt.Println(piscine.DigitLen(-100, 16))
	fmt.Println(piscine.DigitLen(100, -1))
}
