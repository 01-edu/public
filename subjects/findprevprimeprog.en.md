## findprevprime

### Instructions

Write a function that returns the first prime number that is equal or inferior to the `int` passed as parameter.

If there are no primes inferior to the `int` passed as parameter the function should return 0.

### Expected function

```go
func FindPrevPrime(nb int) int {

}
```

### Usage

Here is a possible [program](TODO-LINK) to test your function :

```go
package main

import (
	piscine ".."

	"fmt"
)

func main() {
	fmt.Println(piscine.FindPrevPrime(5))
	fmt.Println(piscine.FindPrevPrime(4))
}
```

And its output :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
5
3
student@ubuntu:~/piscine-go/test$
```
