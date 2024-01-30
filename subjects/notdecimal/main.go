package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Print(piscine.NotDecimal("0.1"))
	fmt.Print(piscine.NotDecimal("174.2"))
	fmt.Print(piscine.NotDecimal("0.1255"))
	fmt.Print(piscine.NotDecimal("1.20525856"))
	fmt.Print(piscine.NotDecimal("-0.0f00d00"))
	fmt.Print(piscine.NotDecimal(""))
	fmt.Print(piscine.NotDecimal("-19.525856"))
	fmt.Print(piscine.NotDecimal("1952"))
}
