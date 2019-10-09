## advancedsortwordarr

### Instructions

Write a function `AdvancedSortWordArr` that sorts a `string` array, based on the function `f` passed in parameter.

### Expected function

```go
func AdvancedSortWordArr(array []string, f func(a, b string) int) {

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

	result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	piscine.AdvancedSortWordArr(result, piscine.Compare)

	fmt.Println(result)
}
```

And its output :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
[1 2 3 A B C a b c]
student@ubuntu:~/piscine-go/test$
```
