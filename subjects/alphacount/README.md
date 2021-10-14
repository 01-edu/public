## alphacount

### Instructions

Write a function that counts the letters of a string and returns the count.

The letters are from the Latin alphabet list only, any other characters, symbols or empty spaces shall not be counted.

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
