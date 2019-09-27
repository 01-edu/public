## activebits

### Instructions

Write a function, `ActiveBits`, that returns the number of active `bits` (bits with the value 1) in the binary representation of an integer number.

### Expected function

```go
func ActiveBits(n int) uint {

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
	nbits := piscine.ActiveBits(7)
	fmt.Println(nbits)
}
```

And its output :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
3
student@ubuntu:~/piscine-go/test$
```
