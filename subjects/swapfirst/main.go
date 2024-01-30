package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.SwapFirst([]int{1, 2, 3, 4}))
	fmt.Println(piscine.SwapFirst([]int{3, 4, 6}))
	fmt.Println(piscine.SwapFirst([]int{1}))
}
