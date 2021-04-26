## split

### Instructions

Write a function that separates the words of a `string` and puts them in a `string` slice.

The separators are the characters of the separator string given in parameter.

### Expected function

```go
func Split(s, sep string) []string {

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
	s := "HelloHAhowHAareHAyou?"
	fmt.Printf("%#v\n", piscine.Split(s, "HA"))
}
```

And its output :

```console
student@ubuntu:~/split/test$ go build
student@ubuntu:~/split/test$ ./test
[]string{"Hello", "how", "are", "you?"}
student@ubuntu:~/split/test$
```
