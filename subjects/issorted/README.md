## issorted

### Instructions

Write a function `IsSorted` that returns `true` if the slice of `int` is sorted, and that returns `false` otherwise.

The function passed in parameter returns a positive `int` if `a` (the first argument) is superior to `b` (the second argument), it returns `0` if they are equal and it returns a negative `int` otherwise.

To do your testing you have to write your own `f` function.

### Expected function

```go
func IsSorted(f func(a, b int) int, a []int) bool {

}
```

### Usage

Here is a possible program to test your function (without `f`):

```go
package main

import (
	"fmt"
	piscine ".."
)

func main() {
	a1 := []int{0, 1, 2, 3, 4, 5}
	a2 := []int{0, 2, 1, 3}

	result1 := piscine.IsSorted(f, a1)
	result2 := piscine.IsSorted(f, a2)

	fmt.Println(result1)
	fmt.Println(result2)
}
```

And its output :

```console
$ go run .
true
false
$
```
