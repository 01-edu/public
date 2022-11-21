## multorsum

### Instructions

Write a function that receives a slice of `int` and an `int` representing the initial value as arguments.

The function should go through the slice and for each `int` check the following restrictions:

- If the `int` is odd, multiply it by the init value. Return the accumulated value after traversing the entire slice.
- If the `int` is even, add it to the init value. Return the accumulated value after traversing the entire slice.
- If the slice is empty return `0`.

### Expected function

```go
func MultOrSum(ints []int, init int) int {

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
	fmt.Println(piscine.MultOrSum([]int{1, 2, 3, 4}, 3))
	fmt.Println(piscine.MultOrSum([]int{1, 2, 3, 4}, 0))
	fmt.Println(piscine.MultOrSum([]int{1, -2, 3, 4}, 0))
}

```

And its output:

```console
$ go run . | cat -e
19$
10$
-2$
```
