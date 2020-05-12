package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var a1, a2, rev []rune
	for _, arg := range os.Args[1:] {
		for _, k := range arg {
			if strings.ContainsRune("aeiouAEIOU", k) {
				a1 = append(a1, k)
			}
		}
	}
	for i := len(a1) - 1; i >= 0; i-- {
		rev = append(rev, a1[i])
	}

	m := 0
	for i, arg := range os.Args[1:] {
		for _, j := range arg {
			if strings.ContainsRune("aeiouAEIOU", j) {
				a2 = append(a2, rev[m])
				m++
			} else {
				a2 = append(a2, j)
			}
		}
		if i != len(os.Args)-1 {
			a2 = append(a2, ' ')
		}
	}
	fmt.Println(string(a2))
}
