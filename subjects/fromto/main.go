package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Print(piscine.FromTo(1, 10))
	fmt.Print(piscine.FromTo(10, 1))
	fmt.Print(piscine.FromTo(10, 10))
	fmt.Print(piscine.FromTo(100, 10))
}
