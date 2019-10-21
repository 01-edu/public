## alphacount

### Instructions

Write a function that counts only the letters of a `string` and that returns that count.
White spaces or any other characters should not be counted.

The letters are only the ones from the latin alphabet.

### Expected function

```go
func AlphaCount(str string) int {

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
	str := "Hello 78 World!    4455 /"
	nb := piscine.AlphaCount(str)
	fmt.Println(nb)
}
```

And its output :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
10
student@ubuntu:~/piscine-go/test$
```
