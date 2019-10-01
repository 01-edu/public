## iterativefactorial

### Instructions

Write an **iterative** function that returns the factorial of the `int` passed as parameter.

Errors (non possible values or overflows) will return `0`.

### Expected function

```go
func IterativeFactorial(nb int) int {

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
	fmt.Println(piscine.IterativeFactorial(arg))
}
```

And its output :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
24
student@ubuntu:~/piscine-go/test$
```
