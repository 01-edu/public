## activebits

### Instructions

Write a function, `ActiveBits`, that returns the number of active `bits` (bits with the value 1) in the binary representation of an integer number.

### Expected function

```go
func ActiveBits(n int) int {

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
	fmt.Println(piscine.ActiveBits(7))
}
```

And its output :

```console
student@ubuntu:~/activebits/test$ go build
student@ubuntu:~/activebits/test$ ./test
3
student@ubuntu:~/activebits/test$
```
