## index

### Instructions

Write a function that behaves like the [`Index`](https://golang.org/pkg/strings/#Index) function.

### Expected function

```go
func Index(s string, toFind string) int {

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
	fmt.Println(piscine.Index("Hello!", "l"))
	fmt.Println(piscine.Index("Salut!", "alu"))
	fmt.Println(piscine.Index("Ola!", "hOl"))
}
```

And its output :

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test
2
1
-1
student@ubuntu:~/[[ROOT]]/test$
```
