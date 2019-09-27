## sqrt

### Instructions

Write a function that returns the square root of the `int` passed as parameter if that square root is a whole number. Otherwise it returns `0`.

### Expected function

```go
func Sqrt(nb int) int {

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
	fmt.Println(piscine.Sqrt(arg1))
	fmt.Println(piscine.Sqrt(arg2))

}
```

And its output :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
2
0
student@ubuntu:~/piscine-go/test$
```
