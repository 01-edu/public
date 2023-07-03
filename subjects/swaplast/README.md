## swaplast

### Instructions

Write a function that takes as an argument a slice of integers and return another slice of integers with the last two elements swapped.

> If the slice contains less than two elements return the same slice.

### Expected function

```go
func SwapLast(slice []int) []int {

}
```

### Usage

Here is a possible program to test your function:

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println(SwapLast([]int{1, 2, 3, 4}))
	fmt.Println(SwapLast([]int{3, 4, 5}))
	fmt.Println(SwapLast([]int{1}))
}

```

And its output:

```console
$ go run .
[1 2 4 3]
[3 5 4]
[1]
```
