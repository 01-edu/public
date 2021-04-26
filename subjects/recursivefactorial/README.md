## recursivefactorial

### Instructions

Write a **recursive** function that returns the factorial of the `int` passed as parameter.

Errors (non possible values or overflows) will return `0`.

`for` is **forbidden** for this exercise.

### Expected function

```go
func RecursiveFactorial(nb int) int {

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
	arg := 4
	fmt.Println(piscine.RecursiveFactorial(arg))
}
```

```console
student@ubuntu:~/recursivefactorial/test$ go build
student@ubuntu:~/recursivefactorial/test$ ./test
24
student@ubuntu:~/recursivefactorial/test$
```
