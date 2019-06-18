## abort

### Instructions

Write a function that returns the value in the middle of 5 five arguments.

This function must have the following signature.

### Expected function

```go
func Abort(a, b, c, d, e int) int {

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
	middle := piscine.Abort(2, 3, 8, 5, 7)
	fmt.Println(middle)
}
```

And its output :

```console
student@ubuntu:~/piscine/test$ go build
student@ubuntu:~/piscine/test$ ./abort
5
student@ubuntu:~/piscine/test$
```
