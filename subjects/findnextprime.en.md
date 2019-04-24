## findnextprime

### Intructions

Write a function that returns the first prime number that is equal or superior to the `int` passed as parameter.

### Expected function

```go
func FindNextPrime(int nb) int {

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
	fmt.Println(piscine.FindNextPrime(5))
	fmt.Println(piscine.FindNextPrime(4))
}
```

And its output :

```console
student@ubuntu:~/piscine/test$ go build
student@ubuntu:~/piscine/test$ ./test
5
5
student@ubuntu:~/piscine/test$
```
