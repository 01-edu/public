## toupper

### Instructions

Write a function that capitalizes each letter of `string`.

### Expected function

```go
func ToUpper(s string) string {

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
	fmt.Println(piscine.ToUpper("Hello! How are you?"))
}
```

And its output :

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test
HELLO! HOW ARE YOU?
student@ubuntu:~/[[ROOT]]/test$
```
