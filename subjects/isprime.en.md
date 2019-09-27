## isprime

### Instructions

Write a function that returns `true` if the `int` passed as parameter is a prime number. Otherwise it returns `false`.

The function must be optimized in order to avoid time-outs with the tester.

### Expected function

```go
func IsPrime(nb int) bool {

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
	fmt.Println(piscine.IsPrime(5))
	fmt.Println(piscine.IsPrime(4))
}
```

And its output :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
true
false
student@ubuntu:~/piscine-go/test$
```
