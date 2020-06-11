## convertbase

### Instructions

Write a function that returns the conversion of a `string` number from one `string` baseFrom to one `string` baseTo.

Only valid bases will be tested.

Negative numbers will not be tested.

### Expected function

```go
func ConvertBase(nbr, baseFrom, baseTo string) string {

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
	result := piscine.ConvertBase("101011", "01", "0123456789")
	fmt.Println(result)
}
```

And its output :

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test
43
student@ubuntu:~/[[ROOT]]/test$
```
