## alphacount

### Instructions

Write a function that counts only the letters of a `string` and that returns that count.

- White spaces or any other characters should not be counted.

The letters are only the ones from the latin alphabet.

### Expected function

```go
func AlphaCount(s string) int {

}
```

### Usage

Here is a possible program to test your function :

```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	s := "Hello 78 World!    4455 /"
	nb := piscine.AlphaCount(s)
	fmt.Println(nb)
}
```

And its output :

```console
$ go run .
10
$
```
