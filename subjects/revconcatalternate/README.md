## revconcatalternate

### Instructions

Write a function `RevConcatAlternate()` that receives two slices of `int` as arguments and returns a new slice with alternated values of each slice in reverse order.
- The input slices can have different lengths.
- The new slice should start with an element of the largest slice first.
- If the slices are of equal length, the new slice should start with an element of the first slice.

### Expected function

```go
func RevConcatAlternate(slice1,slice2 []int) []int {

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
	fmt.Println(piscine.RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(piscine.RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(piscine.RevConcatAlternate([]int{1, 2, 3}, []int{}))
}
```

And its output:

```console
$ go run .
[3 6 2 5 1 4]
[9 3 8 2 7 1 6 5 4]
[3 2 1]
```
