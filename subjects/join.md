## join

### Instructions

Write a function, Join, that returns the elements of a slice strings (arstr) join together in one string. Using the string 'sep' as a separator between each element of the array

The function must have the next signature.

### Expected function

```go
func Join(arstr []string, sep string) string {

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
	arrStr := []string{"hello", "how", "are", "you"}
	joined := student.Join(arrStr, "--")
	fmt.Println(joined)
}
```

And its output :

```console
student@ubuntu:~/student/join$ go build
student@ubuntu:~/student/join$ ./join
hello--how--are--you
student@ubuntu:~/student/join$
```
