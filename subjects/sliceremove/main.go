package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.SliceRemove([]int{1, 2, 3}, 2))
	fmt.Println(piscine.SliceRemove([]int{4, 3}, 4))
	fmt.Println(piscine.SliceRemove([]int{}, 1))

}
