package main

import (
	solutions "./solutions"

	student "./student"
	"github.com/01-edu/z01"
)

func main() {
	var arr []string
	for l := 0; l < 7; l++ {
		a := z01.RandIntBetween(5, 20)
		b := z01.RandASCII()
		str := z01.RandStr(a, b)
		arr = append(arr, str)
	}

	arr = append(arr, " ")

	// example from the subject
	arr = append(arr, "Hello 78 World!    4455 /")

	for i := 0; i < len(arr); i++ {
		z01.Challenge("AlphaCount", student.AlphaCount, solutions.AlphaCount, arr[i])
	}
}
