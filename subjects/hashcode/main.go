package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.HashCode("A"))
	fmt.Println(piscine.HashCode("AB"))
	fmt.Println(piscine.HashCode("BAC"))
	fmt.Println(piscine.HashCode("Hello World"))
}
