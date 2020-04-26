package correct

import "fmt"

func Chunk(a []int, ch int) {
	slice := []int{}
	if ch <= 0 {
		fmt.Println()
		return
	}
	result := make([][]int, 0, len(a)/ch+1)
	for len(a) >= ch {
		slice, a = a[:ch], a[ch:]
		result = append(result, slice)
	}
	if len(a) > 0 {
		result = append(result, a[:len(a)])
	}
	fmt.Println(result)
}
