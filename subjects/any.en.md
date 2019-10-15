## any

### Instructions

Write a function `Any` that returns `true`, for a `string` array :

-   if, when that `string` array is passed through an `f` function, at least one element returns `true`.

### Expected function

```go
func Any(f func(string) bool, arr []string) bool {

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
	tab1 := []string{"Hello", "how", "are", "you"}
	tab2 := []string{"This", "is", "4", "you"}

	result1 := piscine.Any(piscine.IsNumeric, tab1)
	result2 := piscine.Any(piscine.IsNumeric, tab2)

	fmt.Println(result1)
	fmt.Println(result2)
}
```

And its output :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
false
true
student@ubuntu:~/piscine-go/test$
```
