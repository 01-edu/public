## swap

### Instructions

-   Write a function that swaps the contents of two **pointers to an int** (`*int`).

### Expected function

```go
func Swap(a *int, b *int) {

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
	a := 0
	b := 1
	piscine.Swap(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}
```

And its output :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
1
0
student@ubuntu:~/piscine-go/test$
```
