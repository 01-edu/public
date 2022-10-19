## isprime

### Instructions

Write a function that returns `true` if the `int` passed as parameter is a prime number. Otherwise it returns `false`.

The function must be optimized in order to avoid time-outs with the tester.

(We consider that only positive numbers can be prime numbers)

(We also consider that 1 is **not** a prime number)

### Expected function

```go
func IsPrime(nb int) bool {

}
```

### Usage

Here is a possible program to test your function :

```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.IsPrime(5))
	fmt.Println(piscine.IsPrime(4))
}
```

And its output :

```console
$ go run .
true
false
$
```
