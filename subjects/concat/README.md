## concat

### Instructions

Write a function that returns the concatenation of two `string` passed in arguments.

### Expected function

```go
func Concat(str1 string, str2 string) string {

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
	fmt.Println(piscine.Concat("Hello!", " How are you?"))

}
```

And its output :

```console
$ go run .
Hello! How are you?
$
```
