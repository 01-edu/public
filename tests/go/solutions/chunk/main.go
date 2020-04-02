package main

import "fmt"

func Chunk(slice []int, size int) {
	newSlice := []int{}
	if size <= 0 {
		fmt.Println()
		return
	}
	result := make([][]int, 0, len(slice)/size+1)
	for len(slice) >= size {
		newSlice, slice = slice[:size], slice[size:]
		result = append(result, newSlice)
	}
	if len(slice) > 0 {
		result = append(result, slice[:len(slice)])
	}
	fmt.Println(result)
}

func main() {

}
