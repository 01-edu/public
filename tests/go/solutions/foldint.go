package solutions

import (
	"fmt"
)

func FoldInt(f func(int, int) int, arr []int, n int) {
	result := n
	for _, v := range arr {
		result = f(result, v)
	}
	fmt.Println(result)
}
