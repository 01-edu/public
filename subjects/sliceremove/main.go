package main

import (
	"fmt"
)

func main() {
	fmt.Println(SliceRemove([]int{1, 2, 3}, 2))
	fmt.Println(SliceRemove([]int{4, 3}, 4))
	fmt.Println(SliceRemove([]int{}, 1))

}
