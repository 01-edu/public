package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.Itoa(12345))
	fmt.Println(piscine.Itoa(0))
	fmt.Println(piscine.Itoa(-1234))
	fmt.Println(piscine.Itoa(987654321))
}
