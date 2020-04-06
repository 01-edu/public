## isupper

### Instructions

Write a function that returns `true` if the `string` passed in parameter only contains uppercase characters, and that returns `false` otherwise.

### Expected function

```go
func IsUpper(str string) bool {

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
	fmt.Println(piscine.IsUpper("HELLO"))
	fmt.Println(piscine.IsUpper("HELLO!"))

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
