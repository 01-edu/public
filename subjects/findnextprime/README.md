## findnextprime

### Instructions

Write a function that returns the first prime number that is equal or superior to the `int` passed as parameter.

The function must be optimized in order to avoid time-outs with the tester.

(We consider that only positive numbers can be prime numbers)

### Expected function

```go
func FindNextPrime(nb int) int {

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
	fmt.Println(piscine.FindNextPrime(5))
	fmt.Println(piscine.FindNextPrime(4))
}
```

And its output :

```console
$ go run .
5
5
$
```
