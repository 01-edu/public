## swaplast

### Instructions

- Write a function that takes a slice in parameter and swaps the two last elements then returns the new value 
- If the slice contains less than two elements return the same slice.

### Expected function

```go
func SwapLast(slice []int) []int {
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
    fmt.Println(SwapLast([]int{1,2,3,4}))
    fmt.Println(SwapLast([]int{3,4}))
    fmt.Println(SwapLast([]int{1}))
}
```

And its output :

```console
$ go run .
[1,2,4,3]
[4,3]
[1]
```
