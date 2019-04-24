## isprime

### Intructions

Write a function that returns `true` if the `int` passed as parameter is a prime number. Otherwise it returns `false`.

### Expected function

```go
func IsPrime(int nb) bool {

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
student@ubuntu:~/piscine/test$ go build
student@ubuntu:~/piscine/test$ ./test
true
false
student@ubuntu:~/piscine/test$
```
