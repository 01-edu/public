package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	var arr []rune
	var revarr []rune
	vowels := regexp.MustCompile(`[aeiouAEIOU]`)
	for i := 1; i < len(os.Args); i++ {
		for _, k := range os.Args[i] {
			mached := vowels.MatchString(string(k))

			if mached {
				arr = append(arr, rune(k))
			}
		}
	}
	for i := len(arr) - 1; i >= 0; i-- {
		revarr = append(revarr, arr[i])
	}
	arr2 := []rune{}

	m := 0
	for i := 1; i < len(os.Args); i++ {
		for _, j := range os.Args[i] {
			mached := vowels.MatchString(string(j))
			if mached {
				arr2 = append(arr2, rune(revarr[m]))
				m++
			} else {
				arr2 = append(arr2, rune(j))
			}
		}
		if i != len(os.Args)-1 {
			arr2 = append(arr2, ' ')
		}
	}
	fmt.Println(string(arr2))
}
