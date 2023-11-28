package main

import (
	"fmt"
	"piscine"
)

func main() {
	a := []string{"coding", "algorithm", "ascii", "package", "golang"}
	fmt.Printf("%#v\n", piscine.Slice(a, 1))
	fmt.Printf("%#v\n", piscine.Slice(a, 2, 4))
	fmt.Printf("%#v\n", piscine.Slice(a, -3))
	fmt.Printf("%#v\n", piscine.Slice(a, -2, -1))
	fmt.Printf("%#v\n", piscine.Slice(a, 2, 0))
}
