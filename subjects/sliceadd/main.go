package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.SliceAdd([]int{1, 2, 3}, 4))
	fmt.Println(piscine.SliceAdd([]int{}, 4))
}
