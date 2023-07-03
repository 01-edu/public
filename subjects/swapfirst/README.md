## swapfirst

### Instructions

Write a function that takes a slice of integers as an argument and returns another slice of integers with the two first elements swapped.

- If the slice contains less than two elements return the same slice.

### Expected function

```go
func SwapFirst(slice []int) []int {

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
	fmt.Println(SwapFirst([]int{1, 2, 3, 4}))
	fmt.Println(SwapFirst([]int{3, 4, 6}))
	fmt.Println(SwapFirst([]int{1}))
}
```

And its output:

```console
$ go run .
[2 1 3 4]
[4 3 6]
[1]
```
