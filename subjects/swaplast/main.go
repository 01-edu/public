package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.SwapLast([]int{1, 2, 3, 4}))
	fmt.Println(piscine.SwapLast([]int{3, 4, 5}))
	fmt.Println(piscine.SwapLast([]int{1}))
}
