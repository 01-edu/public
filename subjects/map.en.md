## map

### Instructions

Write a function `Map` that, for an `int` array, applies a function of this type `func(int) bool` on each elements of that array and returns an array of all the return values.

### Expected function

```go
func Map(f func(int) bool, arr []int) []bool {

}
```

### Usage

Here is a possible [program](TODO-LINK) to test your function :

```go
package main

import (
	"fmt"
	piscine ".."
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	result := piscine.Map(piscine.IsPrime, arr)
	fmt.Println(result)
}
```

And its output :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
[false true true false true false]
student@ubuntu:~/piscine-go/test$
```
