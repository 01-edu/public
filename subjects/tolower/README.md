## tolower

### Instructions

Write a function that lowercases each letter of `string`.

### Expected function

```go
func ToLower(s string) string {

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
	fmt.Println(piscine.ToLower("Hello! How are you?"))
}
```

And its output :

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test
hello! how are you?
student@ubuntu:~/[[ROOT]]/test$
```
