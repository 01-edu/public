## iterativepower

### Instructions

Write an **iterative** function that returns the power of the `int` passed as parameter.

Negative powers will return `0`. Overflows do **not** have to be dealt with.

### Expected function

```go
func IterativePower(nb int, power int) int {

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
	arg1 := 4
	arg2 := 3
	fmt.Println(piscine.IterativePower(arg1, arg2))
}
```

And its output :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
64
student@ubuntu:~/piscine-go/test$
```
