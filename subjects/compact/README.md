## Compact

### Instructions

Write a function `Compact` that takes a pointer to a slice of `string`s as the argument.
This function must:

- Return the number of elements with [non-zero value](https://tour.golang.org/basics/12).

- Compact, i.e., delete the elements with zero-values in the slice.

### Expected functions

```go
func Compact(ptr *[]string) int {

}
```

### Usage

Here is a possible program to test your function :

```go
package main

import (
	"fmt"

	"piscine"
)

const N = 6

func main() {
	a := make([]string, N)
	a[0] = "a"
	a[2] = "b"
	a[4] = "c"

	for _, v := range a {
		fmt.Println(v)
	}

	fmt.Println("Size after compacting:", piscine.Compact(&a))

	for _, v := range a {
		fmt.Println(v)
	}
}
```

And its output :

```console
$ go run .
a

b

c

Size after compacting: 3
a
b
c
$
```
