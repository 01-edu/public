## capitalize

### Instructions

Write a function that capitalizes the first letter of each word **and** lowercases the rest of each word of a `string`.

A word is a sequence of **alphanumerical** characters.

### Expected function

```go
func Capitalize(s string) string {

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
	fmt.Println(piscine.Capitalize("Hello! How are you? How+are+things+4you?"))
}
```

And its output :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
Hello! How Are You? How+Are+Things+4you?
student@ubuntu:~/piscine-go/test$
```
