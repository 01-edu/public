## isalpha

### Instructions

Write a function that returns `true` if the `string` passed in parameter only contains alphanumerical characters or is empty, and that returns `false` otherwise.

### Expected function

```go
func IsAlpha(s string) bool {

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
	fmt.Println(piscine.IsAlpha("Hello! How are you?"))
	fmt.Println(piscine.IsAlpha("HelloHowareyou"))
	fmt.Println(piscine.IsAlpha("What's this 4?"))
	fmt.Println(piscine.IsAlpha("Whatsthis4"))

}
```

And its output :

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test
false
true
false
true
student@ubuntu:~/[[ROOT]]/test$
```
