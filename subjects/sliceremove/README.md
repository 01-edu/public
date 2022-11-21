## sliceremove

### Instructions

Write a function that takes a slice of integers and an `int` as arguments, if the given `int` exists in the slice then remove it, otherwise return the slice.

- If the slice is empty, return the slice itself.

### Expected function

```go
func SliceRemove(slice []int , num int) []int {

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
	fmt.Println(piscine.SliceRemove([]int{1, 2, 3}, 2))
	fmt.Println(piscine.SliceRemove([]int{4, 3}, 4))
	fmt.Println(piscine.SliceRemove([]int{}, 1))

}
```

And its output:

```console
$ go run .
[1 3]
[3]
[]
```
