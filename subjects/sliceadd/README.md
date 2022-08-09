## sliceadd

### Instructions

- Write a function that takes a slice of integers and int and adds integer in the slice then returns the slice 
- If the slice is empty, return a slice with the new value

### Expected function

```go
func SliceAdd(slice []int , num int) []int {

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
	fmt.Println(SliceAdd([]int{1, 2, 3}, 4))
	fmt.Println(SliceAdd([]int{}, 4))
}
```

And its output :

```console
$ go run .
[1 2 3 4]
[4]
```
