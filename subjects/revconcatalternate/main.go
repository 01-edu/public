package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(piscine.RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(piscine.RevConcatAlternate([]int{1, 2, 3, 9, 8}, []int{4, 5}))
	fmt.Println(piscine.RevConcatAlternate([]int{1, 2, 3}, []int{}))
}
