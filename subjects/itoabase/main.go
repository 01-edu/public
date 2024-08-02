package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.ItoaBase(10, 2))
	fmt.Println(piscine.ItoaBase(255, 16))
	fmt.Println(piscine.ItoaBase(-42, 4))
	fmt.Println(piscine.ItoaBase(123, 10))
	fmt.Println(piscine.ItoaBase(0, 8))
	fmt.Println(piscine.ItoaBase(255, 2))
	fmt.Println(piscine.ItoaBase(-255, 16))
	fmt.Println(piscine.ItoaBase(15, 16))
	fmt.Println(piscine.ItoaBase(10, 4))
	fmt.Println(piscine.ItoaBase(255, 10))
}
