## secondhalf

### Instructions

Write a function `SecondHalf()` that receives a slice of `int` and returns another slice of `int` with the second half of the values.

> If the length is odd, include the middle value in the result.

### Expected function

```go
func SecondHalf(slice []int) []int {

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
	fmt.Println(piscine.SecondHalf([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
	fmt.Println(piscine.SecondHalf([]int{1, 2, 3}))
}
```

And its output:

```console
$ go run . | cat -e
[6 7 8 9 10]
[2 3]
```
