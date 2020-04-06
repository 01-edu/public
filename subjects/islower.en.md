## islower

### Instructions

Write a function that returns `true` if the `string` passed in parameter only contains lowercase characters, and that returns `false` otherwise.

### Expected function

```go
func IsLower(str string) bool {

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
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test
true
false
student@ubuntu:~/[[ROOT]]/test$
```
