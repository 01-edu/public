## map

### Instructions

Write a function `Map` that, for an `int` slice, applies a function of this type `func(int) bool` on each elements of that slice and returns a slice of all the return values.

### Expected function

```go
func Map(f func(int) bool, a []int) []bool {

}
```

### Usage

Here is a possible program to test your function :

```go
package main

import (
	"fmt"
	piscine ".."
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	result := piscine.Map(piscine.IsPrime, a)
	fmt.Println(result)
}
```

And its output :

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test
[false true true false true false]
student@ubuntu:~/[[ROOT]]/test$
```
