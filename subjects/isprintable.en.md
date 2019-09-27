## isprintable

### Instructions

Write a function that returns `true` if the `string` passed in parameter only contains printable characters, and that returns `false` otherwise.

### Expected function

```go
func IsPrintable(str string) bool {

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
	fmt.Println(piscine.IsPrintable("Hello"))
	fmt.Println(piscine.IsPrintable("Hello\n"))

}
```

And its output :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
true
false
student@ubuntu:~/piscine-go/test$
```
