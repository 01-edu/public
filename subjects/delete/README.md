## delete

### Instructions

Write a function that removes the element at a given position of a slice of ints. It should return a new slice with the result. If the position is out of range, it should return the original slice.

### Expected function

```go
func Delete(ints []int, position int) []int {
}
```

### Usage

Here is a possible program to test your function :

```go
package main

import "fmt"

func main() {
	fmt.Println(Delete([]int{1, 2, 3, 4, 5}, 2))
	fmt.Println(Delete([]int{1, 2, 3, 4, 5}, 3))
	fmt.Println(Delete([]int{1, 2, 3, 4, 5}, 1))
}
```

And its output :

```console
$ go run . | cat -e
[1 3 4 5]$
[1 2 4 5]$
[2 3 4 5]$
```
