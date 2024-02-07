package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(piscine.ConcatAlternate([]int{2, 4, 6, 8, 10}, []int{1, 3, 5, 7, 9, 11}))
	fmt.Println(piscine.ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(piscine.ConcatAlternate([]int{1, 2, 3}, []int{}))
}
