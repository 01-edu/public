## sliceadd

### Instructions

Write a function that takes a slice of integers and an `int` as arguments, adds the `int` to the slice and returns it.

- If the slice is empty, return a slice with the new value.

### Expected function

```go
func SliceAdd(slice []int , num int) []int {

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
	fmt.Println(piscine.SliceAdd([]int{1, 2, 3}, 4))
	fmt.Println(piscine.SliceAdd([]int{}, 4))
}
```

And its output:

```console
$ go run .
[1 2 3 4]
[4]
```
