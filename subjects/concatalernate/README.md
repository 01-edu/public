## Concat-Alternate


### Instructions

- Write a function that receives two slices of int as arguments and returns a new one with the result of alternating the values of each.
- The input slices may be of different lengths.
- The new slice should contain the elements of the larger slice first and then the elements of the smaller slice.
- If the slices are of equal length, the new slice should contain the elements of the first slice first and then the elements of the second slice.
- The new slice should be in the same order as the input slices.

### Expected function

```go
func ConcatAlternate(slice1,slice2 []int) []int {
    
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
	fmt.Println(ConcatAlternate([]int{1,2,3},[]int{4,5,6}))
	fmt.Println(ConcatAlternate([]int{1,2,3},[]int{4,5,6,7,8,9}))
	fmt.Println(ConcatAlternate([]int{1,2,3},[]int{}))
}
```

And its output :

```console
$ go run .
[1, 4, 2, 5, 3, 6]
[1, 4, 2, 5, 3, 6, 7, 8, 9]
[1, 2, 3]
```
