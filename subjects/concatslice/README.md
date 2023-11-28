## concatslice

### Instructions

Write a function `ConcatSlice()` that takes two slices of integers as arguments and returns the concatenation of the two slices.

### Expected function

```go
func ConcatSlice(slice1, slice2 []int) []int {

}
```

### Usage

Here is a possible program to test your function:

```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.ConcatSlice([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(piscine.ConcatSlice([]int{}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(piscine.ConcatSlice([]int{1, 2, 3}, []int{}))
}
```

And its output:

```console
$ go run .
[1 2 3 4 5 6]
[4 5 6 7 8 9]
[1 2 3]
```
