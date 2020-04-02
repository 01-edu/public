package solutions

import "fmt"

func Chunk(arr []int, ch int) {
	slice := []int{}
	if ch <= 0 {
		fmt.Println()
		return
	}
	result := make([][]int, 0, len(arr)/ch+1)
	for len(arr) >= ch {
		slice, arr = arr[:ch], arr[ch:]
		result = append(result, slice)
	}
	if len(arr) > 0 {
		result = append(result, arr[:len(arr)])
	}
	fmt.Println(result)
}
