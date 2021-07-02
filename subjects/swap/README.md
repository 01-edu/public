## swap

### Instructions

- Write a function that takes two **pointers to an `int`** (`*int`) and swaps their contents.

### Expected function

```go
func Swap(a *int, b *int) {

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

func main() {
	a := 0
	b := 1
	piscine.Swap(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}
```

And its output :

```console
$ go run .
1
0
$
```

### Notions

- [Pointers](https://golang.org/ref/spec#Pointer_types)
