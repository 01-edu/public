## strlen

### Instructions

-   Write a function that counts the runes of a `string` and that returns that count.

### Expected function

```go
func StrLen(str string) int {

}
```

### Usage

Here is a possible [program](TODO-LINK) to test your function :

```go
package main

import (
	"fmt"
)

func main() {
	str := "Hello World!"
	nb := StrLen(str)
	fmt.Println(nb)
}
```

And its output :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
12
student@ubuntu:~/piscine-go/test$
```
