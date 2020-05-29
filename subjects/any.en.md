## any

### Instructions

Write a function `Any` that returns `true`, for a `string` slice :

- if, when that `string` slice is passed through an `f` function, at least one element returns `true`.

### Expected function

```go
func Any(f func(string) bool, a []string) bool {

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
	a1 := []string{"Hello", "how", "are", "you"}
	a2 := []string{"This", "is", "4", "you"}

	result1 := piscine.Any(piscine.IsNumeric, a1)
	result2 := piscine.Any(piscine.IsNumeric, a2)

	fmt.Println(result1)
	fmt.Println(result2)
}
```

And its output :

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test
false
true
student@ubuntu:~/[[ROOT]]/test$
```
