## bezero

### Instructions

Write a function that takes a slice of integers as an argument and returns the same slice by initializing all the elements to 0.

- If the slice is empty, return an empty slice.

### Expected function

```go
func BeZero(slice []int) []int {

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
	fmt.Println(BeZero([]int{1, 2, 3, 4, 5, 6}))
	fmt.Println(BeZero([]int{}))
}

```

And its output:

```console
$ go run .
[0 0 0 0 0 0]
[]
```
