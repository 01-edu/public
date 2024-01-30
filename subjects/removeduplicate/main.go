package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.RemoveDuplicate([]int{1, 2, 3, 3, 4, 5, 6, 7, 8, 9, 10}))
	fmt.Println(piscine.RemoveDuplicate([]int{1, 1, 2, 2, 3}))
	fmt.Println(piscine.RemoveDuplicate([]int{}))
}
