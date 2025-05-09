package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Print(piscine.ThirdTimeIsACharm("123456789"))
	fmt.Print(piscine.ThirdTimeIsACharm(""))
	fmt.Print(piscine.ThirdTimeIsACharm("a b c d e f"))
	fmt.Print(piscine.ThirdTimeIsACharm("12"))
}
