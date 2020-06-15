package util

import (
	"fmt"

	util2 "../utilDepth2"
)

func LenWrapperU(ss []string) int {
	return util2.LenWrapper(ss)
}

func NotUsed() {
	b := []string{"just", "something"}
	a := len(b)
	for i, v := range b {
		if i == a-1 {
			fmt.Println("Last element", v)
			continue
		}
		fmt.Println("Element", v)
	}
}
