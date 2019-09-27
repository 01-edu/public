## split

### Instructions

Write a function that separates the words of a `string` and puts them in a `string` array.

The separators are the characters of the `charset string` given in parameter.

### Expected function

```go
func Split(str, charset string) []string {

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
	str := "HelloHAhowHAareHAyou?"
	fmt.Println(piscine.Split(str, "HA"))
}
```

And its output :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
[Hello how are you?]
student@ubuntu:~/piscine-go/test$
```
