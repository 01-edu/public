## splitwhitespaces

### Instructions

Write a function that separates the words of a `string` and puts them in a `string` slice.

The separators are spaces, tabs and newlines.

### Expected function

```go
func SplitWhiteSpaces(s string) []string {

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
	fmt.Printf("%#v\n", piscine.SplitWhiteSpaces("Hello how are you?"))
}
```

And its output :

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test
[]string{"Hello", "how", "are", "you?"}
student@ubuntu:~/[[ROOT]]/test$
```
