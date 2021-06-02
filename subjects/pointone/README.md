## pointone

### Instructions

- Write a function that takes a **pointer to an `int`** as argument and gives to this `int` the value of 1.

### Expected function

```go
func PointOne(n *int) {

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
	n := 0
	piscine.PointOne(&n)
	fmt.Println(n)
}
```

And its output :

```console
$ go run .
1
$
```

### Notions

- [Pointers](https://golang.org/ref/spec#Pointer_types)
