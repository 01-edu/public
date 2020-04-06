## sortwordarr

### Instructions

Write a function `SortWordArr` that sorts by `ascii` (in ascending order) a `string` array.

### Expected function

```go
func SortWordArr(array []string) {

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

	result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	piscine.SortWordArr(result)

	fmt.Println(result)
}
```

And its output :

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test
[1 2 3 A B C a b c]
student@ubuntu:~/[[ROOT]]/test$
```
