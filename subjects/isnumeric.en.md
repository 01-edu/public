## isnumeric

### Instructions

Write a function that returns `true` if the `string` passed in parameter only contains numerical characters, and that returns `false` otherwise.

### Expected function

```go
func IsNumeric(s string) bool {

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
	fmt.Println(piscine.IsNumeric("010203"))
	fmt.Println(piscine.IsNumeric("01,02,03"))
}
```

And its output :

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test
true
false
student@ubuntu:~/[[ROOT]]/test$
```
