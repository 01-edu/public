## join

### Instructions

Write a function that returns the concatenation of all the `string` of a table of `string` **separated** by the separator passed in argument.

### Expected function

```go
func Join(strs []string, sep string) string {

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
	toConcat := []string{"Hello!", " How", " are", " you?"}
	fmt.Println(piscine.Join(toConcat, ":"))
}
```

And its output :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
Hello!: How: are: you?
student@ubuntu:~/piscine-go/test$
```
