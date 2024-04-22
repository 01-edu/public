package main

import (
	"fmt"
	"piscine"
)

func main() {
	input1 := []int{2, 3, 1, 1, 4}
	fmt.Println(piscine.CanJump(input1))

	input2 := []int{3, 2, 1, 0, 4}
	fmt.Println(piscine.CanJump(input2))

	input3 := []int{0}
	fmt.Println(piscine.CanJump(input3))
}
