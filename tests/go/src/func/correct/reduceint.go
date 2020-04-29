package correct

import "fmt"

func ReduceInt(f func(int, int) int, a []int) {
	acc := a[0]
	for i := 1; i < len(a); i++ {
		acc = f(acc, a[i])
	}
	fmt.Println(acc)
}
