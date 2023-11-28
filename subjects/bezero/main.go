package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.BeZero([]int{1, 2, 3, 4, 5, 6}))
	fmt.Println(piscine.BeZero([]int{}))
}
