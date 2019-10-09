## strrev

### Instructions

-   Write a function that reverses a `string`.

-   This function will **return** the s `string`.

### Expected function

```go
func StrRev(s string) string {

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
	s := "Hello World!"
	s = piscine.StrRev(s)
	fmt.Println(s)
}
```

And its output :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
!dlroW olleH
student@ubuntu:~/piscine-go/test$
```
