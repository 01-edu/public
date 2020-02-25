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

Here is a possible program to test your function :

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
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test
10
student@ubuntu:~/[[ROOT]]/test$
```
