package main

import "fmt"

func ReduceInt(f func(int, int) int, arr []int) {
	acc := arr[0]
	for i := 1; i < len(arr); i++ {
		acc = f(acc, arr[i])
	}
	fmt.Println(acc)
}

func main() {

}
