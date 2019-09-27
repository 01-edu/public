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

Here is a possible [program](TODO-LINK) to test your function :

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
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
24
student@ubuntu:~/piscine-go/test$
```
