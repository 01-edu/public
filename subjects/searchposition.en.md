## searchposition

### Instructions

Write a function that receives a slice of strings and a `string` and returns a slice of `int` that represents the indexes of the words that contain the `string`.

### Expected function

```go
func SearchPosition(arr []string, str string) []int {

}
```

### Usage

Here is a possible program to test your function:

```go
package main

import (
	"fmt"
	piscine ".."
)

func main() {
	arr := []string{"a.go", "b.go", "abc", "c.golang"}
	fmt.Println(piscine.SearchPosition(arr, ".go"))    
}
```

And its output :

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test
[0 1 3]
student@ubuntu:~/[[ROOT]]/test$
```