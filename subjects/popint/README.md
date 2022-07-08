## popint

### Instructions

Write a function that receives a slice of int and returns a new slice without the last element. If the slice is empty return an empty slice.

### Expected function

```go
func PopInt(ints []int) []int {

}
```

### Usage

Here is a possible program to test your function:

```go
package main

import "fmt"

func main() {
	fmt.Println(PopInt([]int{6, 7, 8, 9}))
}
```

And its output:

```console
$ go run .
[6 7 8]
```
