## fibonacci

### Instructions

Write an **recursive** function that returns the value of fibonacci sequence matching the index passed as parameter.

The first value is at index `0`.

The sequence starts this way: 0, 1, 1, 2, 3 etc...

A negative index will return `-1`.

`for` is **forbidden** for this exercise.

### Expected function

```go
package piscine

func Fibonacci(index int) int {

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
	arg1 := 4
	fmt.Println(piscine.Fibonacci(arg1))
}
```

And its output :

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test
3
student@ubuntu:~/[[ROOT]]/test$
```
