## abort

### Instructions

Write a function that returns the the value in the middle of 5 five arguments.

This function must have the following signature.

### Expected function

```go
func Abort(a, b, c, d, e int) int {
	
}
```

### Usage

Here is a possible [program](TODO-LINK) to test your function :

```go
package main

import (
  "fmt"
  student ".."
)

func main() {
	middle := student.Abort(2, 3, 8, 5, 7)
	fmt.Println(middle)
}
```

And its output :

```console
student@ubuntu:~/student/abort$ go build
student@ubuntu:~/student/abort$ ./abort
5
student@ubuntu:~/student/abort$ 
```
