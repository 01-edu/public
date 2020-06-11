## compare

### Instructions

Write a function that behaves like the [`Compare`](https://golang.org/pkg/strings/#Compare) function.

### Expected function

```go
func Compare(a, b string) int {

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
	fmt.Println(piscine.Compare("Hello!", "Hello!"))
	fmt.Println(piscine.Compare("Salut!", "lut!"))
	fmt.Println(piscine.Compare("Ola!", "Ol"))
}
```

And its output :

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test
0
-1
1
student@ubuntu:~/[[ROOT]]/test$
```
