package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.HalfSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
	fmt.Println(piscine.HalfSlice([]int{1, 2, 3}))
}
