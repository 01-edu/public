package main

import (
	"fmt"
	"piscine"
)

func main() {
	result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	piscine.SortWordArr(result)

	fmt.Println(result)
}
