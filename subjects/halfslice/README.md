## halfslice

### Instructions

Write a function that receives a slice with integers and returns a new one with the first half of the values. If the length is odd, round it up.

### Expected function

```go
func HalfSlice(slice []int) []int {

}
```

### Usage

Here is a possible program to test your function :

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println(HalfSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
	fmt.Println(HalfSlice([]int{1, 2, 3}))
}
```

And its output :

```console
$ go run .
[1,2,3,4,5]
[1,2]
```
