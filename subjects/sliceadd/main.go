package main

import (
	"fmt"
)

func main() {
	fmt.Println(SliceAdd([]int{1, 2, 3}, 4))
	fmt.Println(SliceAdd([]int{}, 4))
}
