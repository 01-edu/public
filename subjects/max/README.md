## max

### Instructions

Write a function, `Max`, that returns the maximum value in a slice of integers. If the slice is empty, returns 0.

### Expected function

```go
func Max(a []int) int {

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
	a := []int{23, 123, 1, 11, 55, 93}
	max := piscine.Max(a)
	fmt.Println(max)
}
```

And its output :

```console
student@ubuntu:~/max/test$ go build
student@ubuntu:~/max/test$ ./test
123
student@ubuntu:~/max/test$
```
