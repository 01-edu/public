package main

import (
	"fmt"
	"github.com/01-edu/z01"
	"strconv"

	student "../../student"
	"strings"
)

func ItoaBase(value, base int) string {
	if base < 2 || base > 16 {
		return ""
	}
	return strings.ToUpper(strconv.FormatInt(int64(value), base))
}

func main() {
	value := z01.RandIntBetween(-1000000, 1000000)
	base := z01.RandIntBetween(2, 16)
	fmt.Println(student.ItoaBase(141895, 11))

	fmt.Println(ItoaBase(value, base))
}
