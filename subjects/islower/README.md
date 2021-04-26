## islower

### Instructions

Write a function that returns `true` if the `string` passed in parameter only contains lowercase characters, and that returns `false` otherwise.

### Expected function

```go
func IsLower(s string) bool {

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
	fmt.Println(piscine.IsLower("hello"))
	fmt.Println(piscine.IsLower("hello!"))

}
```

And its output :

```console
student@ubuntu:~/islower/test$ go build
student@ubuntu:~/islower/test$ ./test
true
false
student@ubuntu:~/islower/test$
```
