## strlen

### Instructions

- Write a function that counts the `runes` of a `string` and that returns that count.

### Expected function

```go
func StrLen(str string) int {

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
	str := "Hello World!"
	nb := piscine.StrLen(str)
	fmt.Println(nb)
}
```

And its output :

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test
12
student@ubuntu:~/[[ROOT]]/test$
```
